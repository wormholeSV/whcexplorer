create database  if not exists wormhole;
use wormhole;
CREATE TABLE `txes` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `tx_id` int(11) NOT NULL,
  `tx_hash` char(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `protocol` enum('fiat','bitcoincash','wormhole') COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'bitcoincash',
  `tx_type` bigint(20) unsigned NOT NULL,
  `ecosystem` enum('production','testsystem') COLLATE utf8mb4_unicode_ci DEFAULT 'production',
  `tx_state` enum('pending','valid','invalid') COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'pending',
  `tx_error_code` int(11) DEFAULT '0',
  `tx_block_height` int(10) unsigned DEFAULT '0',
  `tx_seq_in_block` int(11) DEFAULT '0',
  `block_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_txes_tx_hash` (`tx_hash`),
  KEY `idx_txes_block_time` (`block_time`),
  KEY `idx_txes_tx_id` (`tx_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2316 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
