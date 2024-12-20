// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type GetProductReq struct {
	Id uint32 `json:"id"`
}

type GetProductResp struct {
	Product Product `json:"product"`
}

type ListProductReq struct {
	Page     int32 `json:"page"`
	PageSize int64 `json:"pageSize"`
}

type ListProductResp struct {
	Total    int32     `json:"total"`
	Products []Product `json:"products"`
}

type Product struct {
	Id          uint32    `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Picture     string   `json:"picture"`
	Price       float32  `json:"price"`
	Category    []string `json:"category"`
}

type SearchProductReq struct {
	Query string `json:"query"`
}

type SearchProductResp struct {
	Products []Product `json:"products"`
}
