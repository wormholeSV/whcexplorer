create database  if not exists wormhole;
use wormhole;

CREATE TABLE `addresses_in_txes` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `address` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `property_id` bigint(20) NOT NULL,
  `protocol` enum('fiat','bitcoincash','wormhole') COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'bitcoincash',
  `tx_id` int(11) NOT NULL,
  `address_tx_index` int(11) NOT NULL,
  `address_role` enum('sender','recipient','issuer','participant','payee','seller','buyer','feepayer','payer') COLLATE utf8mb4_unicode_ci NOT NULL,
  `balance_available_credit_debit` double DEFAULT NULL,
  `balance_reserved_credit_debit` double DEFAULT NULL,
  `balance_accepted_credit_debit` double DEFAULT NULL,
  `balance_frozen_credit_debit` double DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_addresses_in_txes_address` (`address`)
) ENGINE=InnoDB AUTO_INCREMENT=6477 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
