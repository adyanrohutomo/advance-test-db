CREATE TABLE `employees` (
  `EmployeeID` int unsigned NOT NULL AUTO_INCREMENT,
  `FirstName` varchar(50) DEFAULT NULL,
  `LastName` varchar(50) DEFAULT NULL,
  `Title` varchar(50) DEFAULT NULL,
  `WorkPhone` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`EmployeeID`)
);