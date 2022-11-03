-- 4.a List of customers located in Irvine city.
SELECT * FROM customers WHERE City = 'Irvin';

-- 4.b List of customers whose order is handled by an employee named Adam Barr.
SELECT * FROM customers c JOIN orders o ON c.CustomerID = o.CustomerID JOIN employees e ON o.EmployeeID = e.EmployeeID WHERE e.FirstName = 'Adam' AND e.LastName = 'Barr';

-- 4.c List of products which are ordered by "Contonso, Ltd" Company.
SELECT * FROM products p JOIN order_details od ON p.ProductID = od.ProductID JOIN orders o ON od.OrderID = o.OrderID JOIN customers c ON o.CustomerID = c.CustomerID WHERE c.CompanyName = 'Contonso, Ltd';

-- 4.d List of transactions (orders) which has "UPS Ground" as shipping method.
SELECT * FROM orders o JOIN shipping_methods sm ON o.ShippingMethodID = sm.ShippingMethodID WHERE sm.ShippingMethod = 'UPS Ground';

-- 4.e List of total cost (including tax and freight charge) for every order sorted by ship date.
SELECT *, od.Quantity * od.UnitPrice - od.Discount + o.FreightCharge + o.Taxes as TotalCost FROM orders o JOIN order_details od ON o.OrderID = od.OrderID 