package model

type Employee struct {
	EmployeeID int64  `gorm:"column:EmployeeID"`
	FirstName  string `gorm:"column:FirstName"`
	LastName   string `gorm:"column:LastName"`
	Title      string `gorm:"column:Title"`
	WorkPhone  string `gorm:"column:WorkPhone"`
}
