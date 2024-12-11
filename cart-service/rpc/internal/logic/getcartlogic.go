package logic

import (
	"context"

	"cart-service/cart"
	"cart-service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCartLogic {
	return &GetCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCartLogic) GetCart(in *cart.GetCartReq) (*cart.GetCartResp, error) {
	// todo: add your logic here and delete this line

	return &cart.GetCartResp{}, nil
}
