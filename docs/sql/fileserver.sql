CREATE TABLE `category` (
`id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
`category_no` varchar(64) NOT NULL COMMENT '分类编号',
`category_title` tinytext NOT NULL COMMENT '分类标题',
`category_desc` text NULL COMMENT '分类描述',
PRIMARY KEY (`id`) 
);

CREATE TABLE `file_info` (
`id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
`file_sha1` char(40) NOT NULL COMMENT '文件hash',
`file_name` varchar(256) NOT NULL COMMENT '文件名',
`file_size` bigint NOT NULL COMMENT '文件大小',
`file_address` text NULL COMMENT '文件存储位置',
`category_no` varchar(64) NULL COMMENT '所属分类',
`create_at` datetime NOT NULL DEFAULT Now() COMMENT '创建时间',
`update_at` datetime NOT NULL DEFAULT Now() ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
`status` tinyint NULL DEFAULT 1 COMMENT '状态(1、可用 2、禁用 3、已删除)',
PRIMARY KEY (`id`) ,
UNIQUE INDEX `uq_file_sha1` (`file_sha1`)
);

CREATE TABLE `user` (
`id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
`user_no` varchar(64) NOT NULL COMMENT '用户编号',
`account` varchar(64) NOT NULL COMMENT '登录账号',
`pwd` varchar(256) NOT NULL COMMENT '登录密码',
`user_name` varchar(64) NULL COMMENT '用户名',
`phone` varchar(64) NULL COMMENT '手机',
`email` varchar(128) NULL COMMENT '邮箱',
`signup_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '注册时间',
`last_active` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后活动时间',
`profile` text NULL,
`status` tinyint NULL DEFAULT 1 COMMENT '状态(1、可用 2、禁用 3、已删除)',
PRIMARY KEY (`id`) ,
UNIQUE INDEX `uq_user_no` (`user_no`),
UNIQUE INDEX `uq_account` (`account`),
UNIQUE INDEX `uq_user_name` (`user_name`)
);



CREATE TABLE `user_file` (
`id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
`user_no` varchar(64) NOT NULL COMMENT '用户编号',
`file_sha1` varchar(255) NOT NULL COMMENT '文件hash',
`file_name` varchar(255) NULL COMMENT '文件名',
`upload_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP  COMMENT '上传时间',
`last_update` datetime NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
`status` int NULL DEFAULT 1 COMMENT '状态(1、可用 2、禁用 3、已删除)',
PRIMARY KEY (`id`) 
);



