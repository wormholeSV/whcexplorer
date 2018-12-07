create database  if not exists wormhole;
use wormhole;
CREATE TABLE `smart_properties` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `protocol` enum('fiat','bitcoincash','wormhole') COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'bitcoincash',
  `property_id` bigint(20) DEFAULT NULL,
  `issuer` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `ecosystem` enum('production','testsystem') COLLATE utf8mb4_unicode_ci DEFAULT 'production',
  `create_tx_id` int(11) NOT NULL,
  `last_tx_id` int(11) NOT NULL,
  `property_name` varchar(256) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `precision` int(11) DEFAULT NULL,
  `prev_property_id` bigint(20) DEFAULT '0',
  `property_service_url` varchar(256) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `property_category` varchar(256) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `property_subcategory` varchar(256) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `registration_data` varchar(5000) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `property_data` json DEFAULT NULL,
  `flags` text COLLATE utf8mb4_unicode_ci,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_smart_properties_property_id` (`property_id`)
) ENGINE=InnoDB AUTO_INCREMENT=398 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
