package logic

import (
	"auth-service/auth"
	"context"
	"database/sql"
	"net/http"
	"user-service/model"

	"user-service/api/biz"
	"user-service/api/internal/svc"
	"user-service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	email := req.Email
	password := req.Password
	userModel := model.NewGomallUserModel(l.svcCtx.Conn)
	user, err := userModel.FindOneByEmail(l.ctx, email)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, biz.NoRegisterError
		default:
			return nil, biz.NewError(http.StatusInternalServerError, err.Error())
		}
	}

	if !biz.CheckPassWord(password, user.Password) {
		return nil, biz.PasswordNotMathError
	}
	res, err := l.svcCtx.AuthRpc.DeliverTokenByRPC(l.ctx, &auth.DeliverTokenReq{UserId: uint32(user.Id)})
	if err != nil {
		return nil, biz.TokenError
	}
	key := l.svcCtx.Config.RedisPrefix + "token:" + res.Token
	val, _ := l.svcCtx.RedisClient.Get(key)
	if val != "" {
		return &types.LoginResp{Token: res.Token}, nil
	}

	// 数据写入redis
	err = l.svcCtx.RedisClient.SetexCtx(context.Background(), "token:"+res.Token, string(user.Id), int(l.svcCtx.Config.RedisConfig.PingTimeout))
	if err != nil {
		return nil, biz.RedisError
	}
	resp = &types.LoginResp{Token: res.Token}
	return
}
