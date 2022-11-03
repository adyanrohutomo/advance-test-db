CREATE TABLE `orders` (
  `OrderID` int unsigned NOT NULL AUTO_INCREMENT,
  `CustomerID` int unsigned NOT NULL,
  `EmployeeID` int unsigned NOT NULL,
  `OrderDate` date NOT NULL,
  `PurchaseOrderNumber` varchar(30) NOT NULL,
  `ShipDate` date DEFAULT NULL,
  `ShippingMethodID` int unsigned DEFAULT NULL,
  `FreightCharge` int DEFAULT NULL,
  `Taxes` int DEFAULT NULL,
  `PaymentReceived` char(1) DEFAULT 0,
  `Comment` varchar(150) DEFAULT NULL,
  PRIMARY KEY (`OrderID`)
  FOREIGN KEY (`CustomerID`) REFERENCES customers(`CustomerID`)
  FOREIGN KEY (`EmployeeID`) REFERENCES employees(`EmployeeID`)
  FOREIGN KEY (`ShippingMethodID`) REFERENCES shipping_methods(`ShippingMethodID`)
);