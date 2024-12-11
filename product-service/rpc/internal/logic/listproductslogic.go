package logic

import (
	"context"
	"encoding/json"

	"product-service/model"
	"product-service/rpc/internal/svc"
	"product-service/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListProductsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListProductsLogic {
	return &ListProductsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListProductsLogic) ListProducts(in *product.ListProductsReq) (*product.ListProductsResp, error) {

	productModel := model.NewProductModel(*l.svcCtx.Conn)
	products, err := productModel.ListProducts(l.ctx, in.Page, in.PageSize)
	if err != nil {
		return nil, err
	}
	var typeProducts []*product.Product

	// 将数据库的model类型的product转换为返回给用户的类型的product
	for _, p := range products {
		var category []string
		if err := json.Unmarshal([]byte(p.Categories.String), &category); err != nil {
			return nil, err
		}
		typeProducts = append(typeProducts, &product.Product{
			Id:          p.Id,
			Name:        p.Name,
			Description: p.Description.String,
			Price:       p.Price,
			Categories:  category,
		})
	}
	resp := &product.ListProductsResp{
		Products: typeProducts,
	}
	return resp, nil
}
