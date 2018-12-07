create database  if not exists wormhole;
use wormhole;

CREATE TABLE if not exists `address_balances` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `address` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `property_id` bigint(20) NOT NULL DEFAULT '0',
  `protocol` enum('fiat','bitcoincash','wormhole') COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'bitcoincash',
  `ecosystem` enum('production','testsystem') COLLATE utf8mb4_unicode_ci DEFAULT 'production',
  `balance_available` double NOT NULL DEFAULT '0',
  `balance_reserved` double NOT NULL DEFAULT '0',
  `balance_accepted` double NOT NULL DEFAULT '0',
  `balance_frozen` double NOT NULL DEFAULT '0',
  `last_tx_id` int(11) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_address_balances_address` (`address`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

