package model

type ShippingMethod struct {
	ShippingMethodID int64  `gorm:"column:ShippingMethodID"`
	ShippingMethod   string `gorm:"column:ShippingMethod"`
}
