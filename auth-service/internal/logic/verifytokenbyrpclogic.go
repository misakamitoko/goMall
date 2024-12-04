package logic

import (
	"context"

	"auth-service/auth"
	"auth-service/biz"
	"auth-service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyTokenByRPCLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyTokenByRPCLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyTokenByRPCLogic {
	return &VerifyTokenByRPCLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VerifyTokenByRPCLogic) VerifyTokenByRPC(in *auth.VerifyTokenReq) (*auth.VerifyResp, error) {
	// todo: add your logic here and delete this line
	token := in.Token
	secretKey := l.svcCtx.Config.Auth.SecretKey
	ok := biz.ParseJwt(token, secretKey)
	return &auth.VerifyResp{Res: ok}, nil
}
