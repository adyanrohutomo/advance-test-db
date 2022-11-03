package model

type Customer struct {
	CustomerID          int64  `gorm:"column:CustomerID"`
	CompanyName         string `gorm:"column:CompanyName"`
	FirstName           string `gorm:"column:FirstName"`
	LastName            string `gorm:"column:LastName"`
	BillingAddress      string `gorm:"column:BillingAddress"`
	City                string `gorm:"column:City"`
	StateOrProvince     string `gorm:"column:StateOrProvince"`
	ZIPCode             string `gorm:"column:ZIPCode"`
	Email               string `gorm:"column:Email"`
	CompanyWebsite      string `gorm:"column:CompanyWebsite"`
	PhoneNumber         string `gorm:"column:PhoneNumber"`
	FaxNumber           string `gorm:"column:FaxNumber"`
	ShipAddress         string `gorm:"column:ShipAddress"`
	ShipCity            string `gorm:"column:ShipCity"`
	ShipStateOrProvince string `gorm:"column:ShipStateOrProvince"`
	ShipZIPCode         string `gorm:"column:ShipZIPCode"`
	ShipPhoneNumber     string `gorm:"column:ShipPhoneNumber"`
}
