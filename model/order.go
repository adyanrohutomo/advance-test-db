package model

import "time"

type Order struct {
	OrderID             int64     `gorm:"column:OrderID"`
	CustomerID          int64     `gorm:"column:CustomerID"`
	EmployeeID          int64     `gorm:"column:EmployeeID"`
	OrderDate           time.Time `gorm:"column:OrderDate"`
	PurchaseOrderNumber string    `gorm:"column:PurchaseOrderNumber"`
	ShipDate            time.Time `gorm:"column:ShipDate"`
	ShippingMethodID    int64     `gorm:"column:ShippingMethodID"`
	FreightCharge       int64     `gorm:"column:FreightCharge"`
	Taxes               int64     `gorm:"column:Taxes"`
	PaymentReceived     string    `gorm:"column:PaymentReceived"`
	Comment             string    `gorm:"column:Comment"`
}
