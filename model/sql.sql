CREATE TABLE `users` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `created_at` datetime default now(),
    `updated_at` datetime default now(),
    `deleted_at` datetime default null,
    `username` varchar(30) NOT NULL COMMENT '账号',
    `password` varchar(100) NOT NULL COMMENT '密码',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `foods` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `created_at` datetime default now(),
    `updated_at` datetime default now(),
    `deleted_at` datetime default null,
    `name` varchar(30),
    `price` decimal(2, 6),
    `type_id` int,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;