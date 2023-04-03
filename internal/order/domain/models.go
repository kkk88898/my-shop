package domain

type OrderDto struct {
	Name       string  `json:"name"`
	Type       int32   `json:"type"`
	Price      float64 `json:"price"`
	Image      string  `json:"image"`
	CreateTime string  `json:"createTime"`
	UpdateTime string  `json:"updateTime"`
}
