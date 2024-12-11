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
	id := uint64(req.Id)
	var val string
	key := l.svcCtx.Config.RedisPrefix + "product:" + string(id)
	// 先查询是否在redis中
	val, err = l.svcCtx.RedisClient.Get(key)
	if val != "" {
		json.Unmarshal([]byte(val), &resp)
		return
	}
	productModel := model.NewProductModel(*l.svcCtx.Conn)
	product, err := productModel.FindOne(context.Background(), id)
	if err != nil {
		return nil, err
	}
	var category []string
	if err := json.Unmarshal([]byte(product.Categories.String), &category); err != nil {
		return nil, err
	}
	resp = &types.GetProductResp{
		Product: types.Product{
			Id:          product.Id,
			Name:        product.Name,
			Description: product.Description.String,
			Picture:     product.Picture.String,
			Price:       product.Price,
			Category:    category,
		},
	}
	// 将数据写入redis
	value, err := json.Marshal(resp)
	l.svcCtx.RedisClient.SetexCtx(context.Background(), key, string(value), 60)
	return
}
