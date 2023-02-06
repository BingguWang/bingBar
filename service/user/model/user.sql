CREATE TABLE `user`
(
    `id`         INT ( 10 ) UNSIGNED NOT NULL AUTO_INCREMENT,
    `created_at` datetime                                DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime                                DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` datetime                                DEFAULT NULL,
    `mobile`     VARCHAR(32) COLLATE utf8mb4_unicode_ci  DEFAULT "" COMMENT '手机号',
    `sex`        INT ( 1 ) UNSIGNED,
    `version`    INT ( 10 ) UNSIGNED DEFAULT 0,
    `auth_key`   VARCHAR(128) COLLATE utf8mb4_unicode_ci DEFAULT "" COMMENT '平台唯一id',
    `password`   VARCHAR(32) COLLATE utf8mb4_unicode_ci  DEFAULT "" COMMENT '密码',
    `nick_name`  VARCHAR(32) COLLATE utf8mb4_unicode_ci  DEFAULT "" COMMENT '昵称',
    `auth_type`  VARCHAR(128) COLLATE utf8mb4_unicode_ci DEFAULT "" COMMENT '平台类型',
    `info`       VARCHAR(256) COLLATE utf8mb4_unicode_ci DEFAULT "" COMMENT 'info',
    PRIMARY KEY (`id`)
) DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;