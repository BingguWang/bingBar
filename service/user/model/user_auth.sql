CREATE TABLE `user_auth`
(
    `id`         INT ( 10 ) UNSIGNED NOT NULL AUTO_INCREMENT,
    `created_at` datetime                                DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime                                DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` datetime                                DEFAULT NULL,
    `version`    INT ( 10 ) UNSIGNED DEFAULT 0,
    `auth_key`   VARCHAR(128) COLLATE utf8mb4_unicode_ci DEFAULT "" COMMENT '平台唯一id',
    `auth_type`  VARCHAR(128) COLLATE utf8mb4_unicode_ci DEFAULT "" COMMENT '平台类型',
    `user_id`    INT ( 10 ) UNSIGNED NOT NULL,
    PRIMARY KEY (`id`),
    KEY `k_t_u_idx` (`auth_key`,`auth_type`)
) DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;

