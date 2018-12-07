create database  if not exists wormhole;
use wormhole;

CREATE TABLE `blocks` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `block_height` bigint(20) NOT NULL,
  `block_hash` char(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `prev_block` char(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `block_time` timestamp NULL DEFAULT NULL,
  `txcount` int(11) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_blocks_block_hash` (`block_hash`),
  KEY `idx_blocks_block_height` (`block_height`)
) ENGINE=InnoDB AUTO_INCREMENT=13737 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

