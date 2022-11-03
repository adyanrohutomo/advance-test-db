package model

type Product struct {
	ProductID   int64  `gorm:"column:ProductID"`
	ProductName string `gorm:"column:ProductName"`
	UnitPrice   int64  `gorm:"column:UnitPrice"`
	InStock     string `gorm:"column:InStock"`
}
