-- MySQL dump 10.13  Distrib 8.0.32, for Linux (x86_64)
--
-- Host: localhost    Database: db_fin
-- ------------------------------------------------------
-- Server version	8.0.32

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `company_procurement_info`
--

DROP TABLE IF EXISTS `company_procurement_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `company_procurement_info` (
  `id` char(10) DEFAULT NULL,
  `supplier_id` char(5) DEFAULT NULL,
  `supplier_contact` char(12) DEFAULT NULL,
  `ordered_product` char(16) DEFAULT NULL,
  `stock_location` char(16) DEFAULT NULL,
  `detail` char(16) DEFAULT NULL,
  `order_unit` char(6) DEFAULT NULL,
  `order_number` double(8,2) DEFAULT NULL,
  `order_unit_price` double(8,2) DEFAULT NULL,
  `restock_date` date DEFAULT NULL,
  KEY `supplier_id` (`supplier_id`),
  CONSTRAINT `company_procurement_info_ibfk_1` FOREIGN KEY (`supplier_id`) REFERENCES `supplier_info` (`supplier_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `company_receivables_info`
--

DROP TABLE IF EXISTS `company_receivables_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `company_receivables_info` (
  `id` char(10) NOT NULL,
  `receivable_sum` double(8,2) DEFAULT NULL,
  `remaining_balance` double(8,2) DEFAULT NULL,
  `customer_name` char(12) DEFAULT NULL,
  `should_get_date` date DEFAULT NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `company_receivables_info_ibfk_1` FOREIGN KEY (`id`) REFERENCES `customer_info` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `customer_info`
--

DROP TABLE IF EXISTS `customer_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `customer_info` (
  `id` char(10) NOT NULL,
  `name` char(12) DEFAULT NULL,
  `phone` char(16) DEFAULT NULL,
  `address` char(30) DEFAULT NULL,
  `age` int DEFAULT NULL,
  `job` char(12) DEFAULT NULL,
  `join_date` date DEFAULT NULL,
  `image` char(10) DEFAULT NULL,
  `permission` char(1) DEFAULT NULL,
  `purchase_status` char(6) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `customer_order_records`
--

DROP TABLE IF EXISTS `customer_order_records`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `customer_order_records` (
  `id` char(10) DEFAULT NULL,
  `order_id` int NOT NULL AUTO_INCREMENT,
  `ordered_product` char(16) DEFAULT NULL,
  `supplier_name` char(16) DEFAULT NULL,
  `unit` char(6) DEFAULT NULL,
  `order_date` date DEFAULT NULL,
  `estimated_submission_date` date DEFAULT NULL,
  `actual_submission_date` date DEFAULT NULL,
  `number` double(8,2) DEFAULT NULL,
  `unit_price` double(8,2) DEFAULT NULL,
  `supplier_id` char(5) DEFAULT NULL,
  PRIMARY KEY (`order_id`),
  KEY `id` (`id`),
  KEY `supplier_id` (`supplier_id`),
  CONSTRAINT `customer_order_records_ibfk_1` FOREIGN KEY (`id`) REFERENCES `customer_info` (`id`),
  CONSTRAINT `customer_order_records_ibfk_2` FOREIGN KEY (`supplier_id`) REFERENCES `supplier_info` (`supplier_id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `supplier_info`
--

DROP TABLE IF EXISTS `supplier_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `supplier_info` (
  `supplier_id` char(5) NOT NULL,
  `supplier_name` char(16) DEFAULT NULL,
  PRIMARY KEY (`supplier_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-12-25 20:08:23
