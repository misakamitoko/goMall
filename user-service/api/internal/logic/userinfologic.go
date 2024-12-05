package logic

import (
	"context"
	"encoding/json"
	"net/http"

	"user-service/api/biz"
	"user-service/api/internal/svc"
	"user-service/api/internal/types"
	"user-service/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line
	userId := l.ctx.Value("userId").(json.Number)

	userModel := model.NewGomallUserModel(l.svcCtx.Conn)
	userId64, err := userId.Int64()
	if err != nil {
		return nil, biz.NewError(http.StatusInternalServerError, err.Error())
	}
	user, err := userModel.FindOne(l.ctx, userId64)
	if err != nil {
		return nil, biz.DbError
	}
	resp = &types.UserInfoResponse{
		Id:     userId64,
		Name:   user.Name,
		Gender: user.Gender,
		Mobile: user.Mobile,
	}
	return
}
