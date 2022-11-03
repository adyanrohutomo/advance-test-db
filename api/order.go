package api

import (
	"net/http"
	"test-database/db"
	"test-database/model"

	"github.com/labstack/echo"
)

func DetailOrder(c echo.Context) error {
	db := db.DbManager()
	orders := []model.OrderDetailResult{}
	db.Table("order_details").
		Select("order_details.OrderDetailID, customers.FirstName, employees.FirstName, shipping_methods.ShippingMethod, products.ProductName, order_details.Quantity, order_details.UnitPrice, order_details.Discount, order_details.Quantity * order_details.UnitPrice - order_details.Discount as SubTotal").
		Joins("JOIN orders ON order_details.OrderID = orders.OrderID").
		Joins("JOIN customers ON orders.CustomerID = customers.CustomerID").
		Joins("JOIN employees ON orders.EmployeeID = employees.EmployeeID").
		Joins("JOIN shipping_methods ON orders.ShippingMethodID = shipping_methods.ShippingMethodID").
		Joins("JOIN products ON order_details.ProductID = products.ProductID").
		Scan(&orders)
	return c.JSON(http.StatusOK, orders)
}
