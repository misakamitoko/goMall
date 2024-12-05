package logic

import (
	"context"
	"time"
	"user-service/api/biz"
	"user-service/model"

	"user-service/api/internal/svc"
	"user-service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	name, password, email, gender, mobile := req.Name, req.Password, req.Email, req.Gender, req.Mobile
	id, err := biz.GenerateId()
	if err != nil {
		return nil, err
	}
	hashedPassword, err := biz.HashPassword(password)
	if err!= nil {
		return nil, err
	}
	userModel := model.NewGomallUserModel(l.svcCtx.Conn)
	_, err = userModel.Insert(l.ctx, &model.GomallUser{
		Id:         id,
		Name:       name,
		Password:   hashedPassword,
		Mobile:     mobile,
		Email:      email,
		Gender:     gender,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return &types.RegisterResp{id}, nil
}
