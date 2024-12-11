package logic

import (
	"context"
	"encoding/json"

	"product-service/api/internal/svc"
	"product-service/api/internal/types"
	"product-service/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListProductsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListProductsLogic {
	return &ListProductsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListProductsLogic) ListProducts(req *types.ListProductReq) (resp *types.ListProductResp, err error) {
	productModel := model.NewProductModel(*l.svcCtx.Conn)
	products, err := productModel.ListProducts(l.ctx, req.Page, req.PageSize)
	if err != nil {
		return nil, err
	}
	var typeProducts []types.Product

	// 将数据库的model类型的product转换为返回给用户的类型的product
	for _, product := range products {
		var category []string
		if err := json.Unmarshal([]byte(product.Categories.String), &category); err != nil {
			return nil, err
		}
		typeProducts = append(typeProducts, types.Product{
			Id:          product.Id,
			Name:        product.Name,
			Description: product.Description.String,
			Price:       product.Price,
			Category:    category,
		})
	}
	resp = &types.ListProductResp{
		Total:    int32(len(products)),
		Products: typeProducts,
	}
	return
}
