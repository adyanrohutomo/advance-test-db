package model

type OrderDetail struct {
	OrderDetailID int64 `gorm:"column:OrderDetailID"`
	OrderID       int64 `gorm:"column:OrderID"`
	ProductID     int64 `gorm:"column:ProductID"`
	Quantity      int64 `gorm:"column:Quantity"`
	UnitPrice     int64 `gorm:"column:UnitPrice"`
	Discount      int64 `gorm:"column:Discount"`
}

type OrderDetailResult struct {
	OrderDetailID  int64  `gorm:"column:OrderDetailID"`
	CustomersName  string `gorm:"column:FirstName"`
	EmployeeName   string `gorm:"column:FirstName"`
	ShippingMethod string `gorm:"column:ShippingMethod"`
	ProductName    string `gorm:"column:ProductName"`
	Quantity       int64  `gorm:"column:Quantity"`
	UnitPrice      int64  `gorm:"column:UnitPrice"`
	Discount       int64  `gorm:"column:Discount"`
	SubTotal       int64  `gorm:"column:SubTotal"`
}
