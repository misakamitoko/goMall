package logic

import (
	"auth-service/auth"
	"context"
	"database/sql"
	"net/http"
	"strconv"
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
	// todo: add your logic here and delete this line
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
	res, err := l.svcCtx.AuthRpc.DeliverTokenByRPC(l.ctx, &auth.DeliverTokenReq{UserId: int32(user.Id)})
	if err != nil {
		return nil, biz.TokenError
	}
	// 数据写入redis
	err = l.svcCtx.RedisClient.SetexCtx(context.Background(), "token:"+res.Token, strconv.FormatInt(user.Id, 16), int(l.svcCtx.Config.RedisConfig.PingTimeout))
	if err != nil {
		return nil, biz.RedisError
	}
	resp = &types.LoginResp{Token: res.Token}
	return
}
