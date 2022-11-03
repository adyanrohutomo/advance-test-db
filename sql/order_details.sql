CREATE TABLE `order_details` (
  `OrderDetailID` int unsigned NOT NULL AUTO_INCREMENT,
  `OrderID` int unsigned NOT NULL,
  `ProductID` int unsigned NOT NULL,
  `Quantity` int NOT NULL,
  `UnitPrice` int NOT NULL,
  `Discount` int DEFAULT NULL,
  PRIMARY KEY (`OrderDetailID`)
  FOREIGN KEY (`OrderID`) REFERENCES orders(`OrderID`)
  FOREIGN KEY (`ProductID`) REFERENCES products(`ProductID`)
);