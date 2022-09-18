CREATE DATABASE `caesar` CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_general_ci';

-- 用户表
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
    `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(30) DEFAULT '' COMMENT '账号',
    `password` VARCHAR(200) DEFAULT '' COMMENT '密码',
    `main_password` VARCHAR(200) DEFAULT '' COMMENT '主密码',
    `email`  VARCHAR(40) DEFAULT '' COMMENT '邮箱',
    `real_name` VARCHAR(30) DEFAULT '' COMMENT '姓名',
    `phone` VARCHAR(15) DEFAULT '' COMMENT '手机',
    `status` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '状态, 0-没激活 1-激活 2-注销',
    `created_at` INT(30) UNSIGNED NOT NULL DEFAULT 0,
    `updated_at` INT(30) UNSIGNED NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- oauth 表
DROP TABLE IF EXISTS `oauth_access_tokens`;
CREATE TABLE `oauth_access_tokens` (
    `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` INT(11) DEFAULT 0 COMMENT '外键: user 表 id',
    `client_ip` CHAR(30) DEFAULT NULL COMMENT '登录 IP',
    `token` VARCHAR(500) DEFAULT NULL COMMENT 'token',
    `revoked` TINYINT(1) DEFAULT 0 COMMENT '是否撤销',
    `expires_at` INT(30) UNSIGNED NULL COMMENT '过期时间',
    `created_at` INT(30) UNSIGNED NOT NULL DEFAULT 0,
    `updated_at` INT(30) UNSIGNED NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`),
    KEY `fr_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;