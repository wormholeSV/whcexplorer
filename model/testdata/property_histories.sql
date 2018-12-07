create database  if not exists wormhole;
use wormhole;

CREATE TABLE `property_histories` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `property_id` bigint(20) NOT NULL,
  `tx_id` bigint(20) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_property_histories_tx_id` (`tx_id`)
) ENGINE=InnoDB AUTO_INCREMENT=882 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
