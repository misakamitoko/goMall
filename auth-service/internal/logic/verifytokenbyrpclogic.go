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
	token := in.Token
	secretKey := l.svcCtx.Config.Jwt.SecretKey
	ok := biz.ParseJwt(token, secretKey)
	return &auth.VerifyResp{Res: ok}, nil
}
