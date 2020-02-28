-- phpMyAdmin SQL Dump
-- version 4.9.0.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Feb 28, 2020 at 08:33 AM
-- Server version: 10.4.6-MariaDB
-- PHP Version: 7.3.9

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `outofbox`
--

-- --------------------------------------------------------

--
-- Table structure for table `itemlist`
--

CREATE TABLE `itemlist` (
  `id` int(11) NOT NULL,
  `Customer name` varchar(50) NOT NULL,
  `item` varchar(100) NOT NULL,
  `value` varchar(30) NOT NULL,
  `order_date` date NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `itemlist`
--

INSERT INTO `itemlist` (`id`, `Customer name`, `item`, `value`, `order_date`) VALUES
(1, 'ABC', 'samsung 7 pro', '30$', '2019-10-17'),
(2, 'DEF', 'nokia1200', '$10', '2019-10-08'),
(3, 'GHI', 'laptop34', '$67', '2019-10-01'),
(4, 'pqw', 'chair', '$12', '2019-10-02'),
(5, 'bgr', 'table', '$23', '2019-10-03'),
(6, 'ert', 'cable', '$18', '2019-10-04'),
(7, 'utr', 'towel', '$2', '2019-10-07'),
(8, 'iop', 'shoes', '$38', '2019-10-06'),
(9, 'hjk', 'Bottle', '$16', '2019-10-09');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `itemlist`
--
ALTER TABLE `itemlist`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `itemlist`
--
ALTER TABLE `itemlist`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
