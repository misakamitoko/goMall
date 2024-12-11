package logic

import (
	"context"

	"cart-service/cart"
	"cart-service/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmptyCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEmptyCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmptyCartLogic {
	return &EmptyCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EmptyCartLogic) EmptyCart(in *cart.EmptyCartReq) (*cart.EmptyCartResp, error) {
	// todo: add your logic here and delete this line

	return &cart.EmptyCartResp{}, nil
}
