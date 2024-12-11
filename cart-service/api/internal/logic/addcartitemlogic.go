package logic

import (
	"context"
	"database/sql/driver"

	"cart-service/api/internal/product"
	"cart-service/api/internal/svc"
	"cart-service/api/internal/types"
	"cart-service/model"
	cartItemModel "cart-service/model/cartItem"
	shoppingCartModel "cart-service/model/shoppingCart"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCartItemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCartItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCartItemLogic {
	return &AddCartItemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCartItemLogic) AddCartItem(req *types.AddItemReq) (resp *types.AddIemResp, err error) {
	// 先查询商品是否存在
	productId := req.Item.Product_id
	_, err = l.svcCtx.ProductClient.GetProduct(context.Background(), &product.GetProductReq{Id: productId})
	if err != nil {
		return nil, err
	}
	// 查询购物车是否存在
	userId := req.User_id
	scModel := shoppingCartModel.NewShoppingCartsModel(*l.svcCtx.Conn)
	shoppingCartresp, err := scModel.FindOneByUserId(l.ctx, uint64(userId))
	var res driver.Result
	// 如果购物车不存在，就创建一个
	if err == model.ErrNotFound {
		res, err = scModel.Insert(l.ctx, &shoppingCartModel.ShoppingCarts{
			UserId: uint64(userId),
		})
		if err != nil {
			return nil, err
		}
	}
	// 插入购物车数据
	model := cartItemModel.NewCartItemsModel(*l.svcCtx.Conn)
	var cartId int64
	if res != nil {
		cartId, err = res.LastInsertId()
		if err != nil {
			return nil, err
		}
	} else {
		cartId = shoppingCartresp.Id
	}
	_, err = model.Insert(l.ctx, &cartItemModel.CartItems{
		CartId:    cartId,
		ProductId: productId,
		Quantity:  int64(req.Item.Quantity),
	})
	if err != nil {
		return nil, err
	}
	return
}
