package logic

import (
	"context"
	"encoding/json"

	"product-service/api/internal/svc"
	"product-service/api/internal/types"
	"product-service/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductLogic {
	return &GetProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProductLogic) GetProduct(req *types.GetProductReq) (resp *types.GetProductResp, err error) {
	productModel := model.NewProductModel(*l.svcCtx.Conn)
	product, err := productModel.FindOne(context.Background(), uint64(req.Id))
	if err != nil {
		return nil, err
	}
	var category []string
	if err := json.Unmarshal([]byte(product.Categories.String), &category); err != nil {
		return nil, err
	}
	resp = &types.GetProductResp{
		Product: types.Product{
			Id:          int64(product.Id),
			Name:        product.Name,
			Description: product.Description.String,
			Picture:     product.Picture.String,
			Price:       product.Price,
			Category:    category,
		},
	}
	return
}
