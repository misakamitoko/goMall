package logic

import (
	"context"

	"cart-service/api/internal/svc"
	"cart-service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmptyCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmptyCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmptyCartLogic {
	return &EmptyCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmptyCartLogic) EmptyCart(req *types.EmptyCartReq) (resp *types.EmptyCartResp, err error) {
	// todo: add your logic here and delete this line

	return
}
