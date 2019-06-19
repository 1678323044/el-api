package main

/* 用户信息 */
type UserData struct {
	Username string `json:"username"`
}

type User struct {
	Code int      `json:"code"`
	Data UserData `json:"user"`
}

/* 地址信息 */
type AddressData struct {
	Addr string      `json:"title" bson:"address"`
	City string      `json:"city"`
}

type Address struct {
	Code  int         `json:"code"`
	Data  AddressData `json:"address"`
}

/* 商铺列表 */
type Support struct {
	IconName string  `json:"icon_name" bson:"icon_name"`
}

type DeliveryModeT struct {
	Text  string  `json:"text"`
}

type ShopsData struct {
	ImagePath      string     `bson:"image_path" json:"image_path"`
	Name           string     `json:"name"`
	Supports       []Support  `json:"supports"`
	Rating         float64    `json:"rating"`
	RecentOrderNum int        `json:"recent_order_num" bson:"recent_order_num"`
	DeliveryMode   DeliveryModeT  `json:"delivery_mode" bson:"delivery_mode"`
	FloatMinimumOrderAmount int   `json:"float_minimum_order_amount" bson:"float_minimum_order_amount"`
	FloatDeliveryFee int          `json:"float_delivery_fee" bson:"float_delivery_fee"`
}

type Shops struct {
	Code  int   `json:"code"`
	Data  []ShopsData `json:"data"`
}
