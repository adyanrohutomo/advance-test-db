CREATE TABLE `products` (
  `ProductID` int unsigned NOT NULL AUTO_INCREMENT,
  `ProductName` varchar(50) DEFAULT NULL,
  `UnitPrice` int DEFAULT NULL,
  `InStock` char(1) DEFAULT 1,
  PRIMARY KEY (`ProductID`)
);