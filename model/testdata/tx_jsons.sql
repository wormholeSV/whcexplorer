create database  if not exists wormhole;
use wormhole;

CREATE TABLE `tx_jsons` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `tx_id` int(11) NOT NULL,
  `protocol` enum('fiat','bitcoincash','wormhole') COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'bitcoincash',
  `tx_data` json NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_tx_jsons_tx_id` (`tx_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2316 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
