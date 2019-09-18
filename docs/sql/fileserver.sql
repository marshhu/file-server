-- ----------------------------
-- 分类表
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
`id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
`category_no` varchar(64) NOT NULL COMMENT '分类编号',
`category_title` tinytext NOT NULL COMMENT '分类标题',
`category_desc` text NULL COMMENT '分类描述',
`parent_category` varchar(64) NULL COMMENT '父级分类编号',
`category_path` text NULL COMMENT '分类路径',
PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_croatian_ci;

-- ----------------------------
-- 文件信息表
-- ----------------------------
DROP TABLE IF EXISTS `file_info`;
CREATE TABLE `file_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `file_sha1` char(40) COLLATE utf8_croatian_ci NOT NULL COMMENT '文件hash',
  `file_name` varchar(256) COLLATE utf8_croatian_ci NOT NULL COMMENT '文件名',
  `file_size` bigint(20) NOT NULL COMMENT '文件大小',
  `file_address` text COLLATE utf8_croatian_ci COMMENT '文件存储位置',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `status` tinyint(4) DEFAULT '1' COMMENT '状态(1、可用 2、禁用 3、已删除)',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_file_sha1` (`file_sha1`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_croatian_ci;

-- ----------------------------
-- 用户表
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_no` varchar(64) COLLATE utf8_croatian_ci NOT NULL COMMENT '用户编号',
  `user_name` varchar(64) COLLATE utf8_croatian_ci DEFAULT NULL COMMENT '用户名',
  `password` varchar(256) COLLATE utf8_croatian_ci NOT NULL COMMENT '登录密码',
  `phone` varchar(64) COLLATE utf8_croatian_ci DEFAULT NULL COMMENT '手机',
  `email` varchar(128) COLLATE utf8_croatian_ci DEFAULT NULL COMMENT '邮箱',
  `signup_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '注册时间',
  `last_active` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后活动时间',
  `profile` text COLLATE utf8_croatian_ci,
  `status` tinyint(4) DEFAULT '1' COMMENT '状态(1、可用 2、禁用 3、已删除)',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_user_no` (`user_no`),
  UNIQUE KEY `uq_user_name` (`user_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_croatian_ci;

-- ----------------------------
-- 用户文件关系表
-- ----------------------------
DROP TABLE IF EXISTS `user_file`;
CREATE TABLE `user_file` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_no` varchar(64) COLLATE utf8_croatian_ci NOT NULL COMMENT '用户编号',
  `file_sha1` varchar(255) COLLATE utf8_croatian_ci NOT NULL COMMENT '文件hash',
  `file_name` varchar(255) COLLATE utf8_croatian_ci DEFAULT NULL COMMENT '文件名',
  `upload_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '上传时间',
  `last_update` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
  `status` int(11) DEFAULT '1' COMMENT '状态(1、可用 2、禁用 3、已删除)',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_croatian_ci;

-- ----------------------------
-- 文件分类关系表
-- ----------------------------
DROP TABLE IF EXISTS `file_category`;
CREATE TABLE `file_category` (
`id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
`file_sha1` char(40) NOT NULL COMMENT '文件hash',
`category_no` varchar(64) NOT NULL COMMENT '分类编号'
);




