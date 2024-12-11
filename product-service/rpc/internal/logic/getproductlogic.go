package logic

import (
	"context"
	"encoding/json"

	"product-service/model"
	"product-service/rpc/internal/svc"
	"product-service/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductLogic {
	return &GetProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductLogic) GetProduct(in *product.GetProductReq) (*product.GetProductResp, error) {
	id := uint64(in.Id)
	var val string
	var resp product.GetProductResp
	key := l.svcCtx.Config.RedisPrefix + "product:" + string(id)
	// 先查询是否在redis中
	val, err := l.svcCtx.RedisClient.Get(key)
	if val != "" {
		json.Unmarshal([]byte(val), &resp)
		return &resp, nil
	}
	productModel := model.NewProductModel(*l.svcCtx.Conn)
	productM, err := productModel.FindOne(context.Background(), id)
	if err != nil {
		return nil, err
	}
	var category []string
	if err := json.Unmarshal([]byte(productM.Categories.String), &category); err != nil {
		return nil, err
	}
	resp = product.GetProductResp{
		Product: &product.Product{
			Id:          uint32(productM.Id),
			Name:        productM.Name,
			Description: productM.Description.String,
			Picture:     productM.Picture.String,
			Price:       float32(productM.Price),
			Categories:  category,
		},
	}
	// 将数据写入redis
	value, err := json.Marshal(resp)
	l.svcCtx.RedisClient.SetexCtx(context.Background(), key, string(value), 60)
	return &resp, nil
}
