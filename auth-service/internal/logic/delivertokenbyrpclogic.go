package logic

import (
	"context"
	"time"

	"auth-service/auth"
	"auth-service/biz"
	"auth-service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeliverTokenByRPCLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeliverTokenByRPCLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeliverTokenByRPCLogic {
	return &DeliverTokenByRPCLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeliverTokenByRPCLogic) DeliverTokenByRPC(in *auth.DeliverTokenReq) (*auth.DeliveryResp, error) {
	// todo: add your logic here and delete this line
	userId := in.UserId
	secretKey := l.svcCtx.Config.Auth.SecretKey
	expire := l.svcCtx.Config.Auth.Expire
	token, err := biz.GetJwtToken(secretKey, time.Now().Unix(), expire, userId)
	if err != nil {
		l.Logger.Error("get jwt token failed", err)
		return nil, err
	}
	return &auth.DeliveryResp{Token: token}, nil
}
