CREATE
DATABASE `search_nova`
/*!40100 DEFAULT CHARACTER SET utf8 */;USE
`search_nova`;
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`         bigint(11) unsigned NOT NULL AUTO_INCREMENT,
    `username`   char(32)  NOT NULL DEFAULT '',
    `password`   char(32)  NOT NULL DEFAULT '',
    `expire`     timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `token`      char(128) NOT NULL DEFAULT '',
    `role`       int(11) NOT NULL DEFAULT '0',
    `created_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `username_idx` (`username`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;
CREATE TABLE `page`
(
    `id`         bigint(11) unsigned NOT NULL AUTO_INCREMENT,
    `url`        varchar(512) NOT NULL DEFAULT '',
    `title`      varchar(512) NOT NULL DEFAULT '',
    `describe`   varchar(512) NOT NULL DEFAULT '',
    `keywords`   varchar(255) NOT NULL DEFAULT '',
    `status` int(11) NOT NULL DEFAULT 0,
    `created_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `url_idx` (`url`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;