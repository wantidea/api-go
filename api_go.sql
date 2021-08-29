/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 80012
 Source Host           : localhost:3306
 Source Schema         : api_go

 Target Server Type    : MySQL
 Target Server Version : 80012
 File Encoding         : 65001

 Date: 25/08/2021 10:22:25
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for blog_article
-- ----------------------------
DROP TABLE IF EXISTS `blog_article`;
CREATE TABLE `blog_article`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '文章ID',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '标题',
  `category_id` int(11) NOT NULL DEFAULT 0 COMMENT '分类ID',
  `banner_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '首图',
  `editor_type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '编辑器类型 0-普通富文本 1-Markdown',
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '内容',
  `content_md` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '内容 md',
  `look_total` int(11) NOT NULL DEFAULT 0 COMMENT '浏览次数',
  `created_user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '数据创建人',
  `created_at` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updated_at` int(11) NOT NULL DEFAULT 0 COMMENT '修改时间',
  `deleted_at` int(11) NOT NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 38 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '博客文章表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of blog_article
-- ----------------------------
INSERT INTO `blog_article` VALUES (38, '猪小弟博客', 13, 0, 1, '<h2>开始写作吧，少年！</h2>\n', '## 开始写作吧，少年！', 1, 37, 1629857469, 1629858090, 0);

-- ----------------------------
-- Table structure for blog_article_category
-- ----------------------------
DROP TABLE IF EXISTS `blog_article_category`;
CREATE TABLE `blog_article_category`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '类别ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '类别',
  `created_user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '数据创建人',
  `created_at` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updated_at` int(11) NOT NULL DEFAULT 0 COMMENT '修改时间',
  `deleted_at` int(11) NOT NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '博客文章分类表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of blog_article_category
-- ----------------------------
INSERT INTO `blog_article_category` VALUES (13, '笔记', 37, 1629857361, 1629857361, 0);

-- ----------------------------
-- Table structure for blog_setting
-- ----------------------------
DROP TABLE IF EXISTS `blog_setting`;
CREATE TABLE `blog_setting`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '设置ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '设置名称',
  `value` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '设置值',
  `updated_at` int(11) NOT NULL DEFAULT 0 COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '博客设置表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of blog_setting
-- ----------------------------
INSERT INTO `blog_setting` VALUES (1, 'logo', 'runtime/upload/images/1a475874aeda2b32e7ce476a32c5224a.png', 1629857931);
INSERT INTO `blog_setting` VALUES (2, 'name', '猪小弟', 1628039307);
INSERT INTO `blog_setting` VALUES (3, 'title', '某不知名程序猿', 1618794954);

-- ----------------------------
-- Table structure for rabbitmq_consume_error
-- ----------------------------
DROP TABLE IF EXISTS `rabbitmq_consume_error`;
CREATE TABLE `rabbitmq_consume_error`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `exchange_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '交换机名称',
  `exchange_type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '交换机类型',
  `queue_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '队列名称',
  `routing_key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '路由键',
  `message_content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '消息内容',
  `error_content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '报错信息',
  `ack` tinyint(1) NOT NULL DEFAULT 0 COMMENT 'Ack 确认 0-未确定 1-已确认',
  `created_at` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updated_at` int(11) NOT NULL DEFAULT 0 COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'RabbitMQ 错误日志表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of rabbitmq_consume_error
-- ----------------------------

-- ----------------------------
-- Table structure for system_admin
-- ----------------------------
DROP TABLE IF EXISTS `system_admin`;
CREATE TABLE `system_admin`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `role_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '角色ID',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '密码',
  `avatar_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '头像图片ID',
  `created_user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '数据创建人',
  `created_at` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updated_at` int(11) NOT NULL DEFAULT 0 COMMENT '修改时间',
  `deleted_at` int(11) NOT NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `name`(`name`) USING BTREE COMMENT '用户名'
) ENGINE = InnoDB AUTO_INCREMENT = 37 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户信息表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of system_admin
-- ----------------------------
INSERT INTO `system_admin` VALUES (37, 'admin', 7, '25f9e794323b453885f5181f1b624d0b', 3, 24, 1629857144, 1629857980, 0);
INSERT INTO `system_admin` VALUES (38, 'blog', 8, '25f9e794323b453885f5181f1b624d0b', 4, 37, 1629857237, 1629857986, 0);

-- ----------------------------
-- Table structure for system_log_operation
-- ----------------------------
DROP TABLE IF EXISTS `system_log_operation`;
CREATE TABLE `system_log_operation`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '日志ID',
  `method` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '请求类型',
  `uri` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '接口',
  `request_msg` varchar(10000) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '请求信息',
  `return_msg` varchar(10000) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '返回信息',
  `ip` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '客户端IP',
  `created_user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '数据创建人',
  `create_time` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统请求日志表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of system_log_operation
-- ----------------------------

-- ----------------------------
-- Table structure for system_menu
-- ----------------------------
DROP TABLE IF EXISTS `system_menu`;
CREATE TABLE `system_menu`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
  `parent_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '父级菜单ID',
  `level` tinyint(4) NOT NULL DEFAULT 0 COMMENT '菜单级别',
  `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '菜单名称',
  `title` varchar(50) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '菜单标题',
  `path` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '菜单接口',
  `redirect` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '重定向',
  `component` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '组件',
  `icon` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '图标',
  `sort` int(11) NOT NULL DEFAULT 255 COMMENT '菜单排序',
  `role_list` varchar(1000) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '角色 ID 组 逗号分隔 如：1,3',
  `is_hidden` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否显示 0-不显示 1-显示',
  `is_auth` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否认证 0-不启用认证 1-启用认证',
  `created_user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '数据创建人',
  `created_at` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updated_at` int(11) NOT NULL DEFAULT 0 COMMENT '修改时间',
  `deleted_at` int(11) NOT NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `parent_id`(`parent_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 50 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统菜单表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of system_menu
-- ----------------------------
INSERT INTO `system_menu` VALUES (1, 0, 0, '首页', '首页', '/', '/dashboard', 'layout/index.vue', 'el-icon-s-grid', 1, '', 0, 0, 0, 0, 1627531167, 0);
INSERT INTO `system_menu` VALUES (2, 1, 1, '仪表盘', '仪表盘', '/dashboard', '', 'views/dashboard/index.vue', 'el-icon-s-data', 255, '', 0, 0, 0, 0, 1627531298, 0);
INSERT INTO `system_menu` VALUES (3, 0, 0, 'system', '系统', '/system', '/system/admin/list', 'layout/index.vue', 'el-icon-setting', 5, '7', 0, 0, 0, 0, 1627539275, 0);
INSERT INTO `system_menu` VALUES (4, 3, 1, 'system/admin', '系统管理员', '/system/admin', '/system/admin/list', 'layout/components/RouterViewBox', 'el-icon-s-custom', 255, '', 0, 0, 24, 1617852681, 1627531334, 0);
INSERT INTO `system_menu` VALUES (5, 4, 2, 'admin/add', '添加管理员', '/system/admin/add', '', 'views/system/admin/add', 'el-icon-sort-down', 257, '', 0, 0, 24, 1617853908, 1627540790, 0);
INSERT INTO `system_menu` VALUES (6, 4, 2, 'admin/upd', '编辑管理员', '/system/admin/upd', '', 'views/system/admin/upd', 'el-icon-sort-down', 258, '', 1, 0, 24, 1617854143, 1627540797, 0);
INSERT INTO `system_menu` VALUES (7, 4, 2, 'admin/rePwd', '重置密码', '/system/admin/rePwd', '', 'views/system/admin/rePwd', 'el-icon-sort-down', 259, '', 1, 0, 24, 1617854178, 1627540800, 0);
INSERT INTO `system_menu` VALUES (8, 3, 1, 'menu/index', '菜单管理', '/system/menu', '/system/menu/list', 'layout/components/RouterViewBox', 'el-icon-menu', 258, '', 0, 0, 24, 1617854226, 1627540780, 0);
INSERT INTO `system_menu` VALUES (9, 8, 2, 'menu/add', '添加菜单', '/system/menu/add', '', 'views/system/menu/add', 'el-icon-sort-down', 256, '', 0, 0, 24, 1617854287, 1627540835, 0);
INSERT INTO `system_menu` VALUES (10, 8, 2, 'menu/upd', '编辑菜单', '/system/menu/upd', '', 'views/system/menu/upd', 'el-icon-sort-down', 257, '', 1, 0, 24, 1617854318, 1627540841, 0);
INSERT INTO `system_menu` VALUES (11, 3, 1, 'role', '角色管理', '/system/role', '/system/role/list', 'layout/components/RouterViewBox', 'el-icon-user-solid', 256, '', 0, 0, 24, 1617854345, 1627540735, 0);
INSERT INTO `system_menu` VALUES (12, 11, 2, 'role/add', '添加角色', '/system/role/add', '', 'views/system/role/add', 'el-icon-sort-down', 257, '', 0, 0, 24, 1617854544, 1627540810, 0);
INSERT INTO `system_menu` VALUES (13, 11, 2, 'role/upd', '编辑角色', '/system/role/upd', '', 'views/system/role/upd', 'el-icon-sort-down', 258, '', 1, 0, 24, 1617854576, 1627540812, 0);
INSERT INTO `system_menu` VALUES (14, 3, 1, 'route/index', '权限管理', '/system/route', '/system/route/list', 'layout/components/RouterViewBox', 'el-icon-help', 257, '', 1, 0, 24, 1617860335, 1628045772, 0);
INSERT INTO `system_menu` VALUES (15, 14, 2, 'route/add', '添加权限', '/system/route/add', '', 'views/system/route/add', 'el-icon-sort-down', 257, '', 0, 0, 24, 1617860376, 1627540824, 0);
INSERT INTO `system_menu` VALUES (16, 14, 2, 'route/upd', '编辑权限', '/system/route/upd', '', 'views/system/route/upd', 'el-icon-sort-down', 258, '', 1, 0, 24, 1617860413, 1627540827, 0);
INSERT INTO `system_menu` VALUES (17, 8, 2, 'menu/list', '菜单列表', '/system/menu/list', '', 'views/system/menu/list', '', 255, '', 0, 0, 24, 1617854226, 1627368233, 1627530835);
INSERT INTO `system_menu` VALUES (18, 11, 2, '角色列表', '角色列表', '/system/role/list', '', 'views/system/role/list', 'el-icon-sort-down', 256, '', 0, 0, 24, 1617875912, 1627540807, 0);
INSERT INTO `system_menu` VALUES (19, 14, 2, '权限列表', '权限列表', '/system/route/list', '', 'views/system/route/list', 'el-icon-sort-down', 256, '', 0, 0, 24, 1617876676, 1627540817, 0);
INSERT INTO `system_menu` VALUES (20, 4, 2, '管理员列表', '管理员列表', '/system/admin/list', '', 'views/system/admin/list', 'el-icon-sort-down', 256, '', 0, 0, 24, 1617876789, 1627540685, 0);
INSERT INTO `system_menu` VALUES (21, 8, 2, '树形菜单', '树形菜单', '/system/menu/treeList', '', 'views/system/menu/treeList', 'el-icon-sort-down', 255, '', 0, 0, 24, 1617930124, 1627538268, 0);
INSERT INTO `system_menu` VALUES (22, 0, 0, '博客', '博客', '/blog', '/blog/article/list', 'layout/index.vue', 'el-icon-notebook-2', 3, '7,8', 0, 0, 24, 1617931573, 1627539314, 0);
INSERT INTO `system_menu` VALUES (23, 22, 1, '文章管理', '文章管理', '/blog/article', '/blog/article/list', 'layout/components/RouterViewBox', 'el-icon-folder-opened', 252, '', 0, 0, 24, 1617931757, 1627539522, 0);
INSERT INTO `system_menu` VALUES (24, 23, 2, '文章列表', '文章列表', '/blog/article/list', '', 'views/blog/article/list', 'el-icon-sort-down', 256, '', 0, 0, 24, 1617931827, 1627540534, 0);
INSERT INTO `system_menu` VALUES (25, 23, 2, '添加文章', '添加文章', '/blog/article/add', '', 'views/blog/article/add', 'el-icon-sort-down', 257, '', 0, 0, 24, 1617931866, 1627540538, 0);
INSERT INTO `system_menu` VALUES (26, 23, 2, '编辑文章', '编辑文章', '/blog/article/upd', '', 'views/blog/article/upd', 'el-icon-sort-down', 259, '', 1, 0, 24, 1617931921, 1627540548, 0);
INSERT INTO `system_menu` VALUES (27, 23, 2, '文章详情', '文章详情', '/blog/article/item', '', 'views/blog/article/item', 'el-icon-sort-down', 261, '', 1, 0, 24, 1617931952, 1629440625, 0);
INSERT INTO `system_menu` VALUES (28, 22, 1, '评论管理', '评论管理', '/blog/comment', '/blog/comment/list', 'layout/components/RouterViewBox', 'el-icon-s-comment', 254, '', 1, 0, 24, 1617932104, 1628045741, 0);
INSERT INTO `system_menu` VALUES (29, 28, 2, '添加评论', '添加评论', '/blog/comment/add', '', 'views/blog/comment/add', 'el-icon-sort-down', 257, '', 0, 0, 24, 1617932173, 1627540615, 0);
INSERT INTO `system_menu` VALUES (30, 28, 2, '评论列表', '评论列表', '/blog/comment/list', '', 'views/blog/comment/add', 'el-icon-sort-down', 256, '', 0, 0, 24, 1617932218, 1627540600, 0);
INSERT INTO `system_menu` VALUES (31, 28, 2, '编辑评论', '编辑评论', '/blog/comment/upd', '', 'views/blog/comment/upd', 'el-icon-sort-down', 258, '', 1, 0, 24, 1617932254, 1627540618, 0);
INSERT INTO `system_menu` VALUES (32, 22, 1, '收藏管理', '收藏管理', '/blog/collection', '/blog/collection/list', 'layout/components/RouterViewBox', 'el-icon-star-off', 254, '', 1, 0, 24, 1617932314, 1628045722, 0);
INSERT INTO `system_menu` VALUES (33, 32, 2, '收藏列表', '收藏列表', '/blog/collection/list', '', 'views/blog/collection/list', 'el-icon-sort-down', 255, '', 0, 0, 24, 1617932361, 1627538293, 0);
INSERT INTO `system_menu` VALUES (34, 32, 2, '添加收藏', '添加收藏', '/blog/collection/add', '', 'views/blog/collection/add', 'el-icon-sort-down', 256, '', 0, 0, 24, 1617932395, 1627540626, 0);
INSERT INTO `system_menu` VALUES (35, 32, 2, '编辑收藏', '编辑收藏', '/blog/collection/upd', '', 'views/blog/collection/upd', 'el-icon-sort-down', 257, '', 1, 0, 24, 1617932436, 1627540630, 0);
INSERT INTO `system_menu` VALUES (36, 22, 1, '博客设置', '博客设置', '/blog/setting', '', 'layout/components/RouterViewBox', 'el-icon-s-tools', 255, '', 0, 0, 24, 1618192969, 1627538480, 0);
INSERT INTO `system_menu` VALUES (37, 36, 2, 'Logo', 'Logo', '/blog/setting/updLogo', '', 'views/blog/setting/updLogo', 'el-icon-sort-down', 254, '', 0, 0, 24, 1618193080, 1627540646, 0);
INSERT INTO `system_menu` VALUES (38, 36, 2, '博客名称', '博客名称', '/blog/setting/updName', '', 'views/blog/setting/updName', 'el-icon-sort-down', 255, '', 0, 0, 24, 1618195389, 1627538304, 0);
INSERT INTO `system_menu` VALUES (39, 36, 2, '博客介绍', '博客介绍', '/blog/setting/updTitle', '', 'views/blog/setting/updTitle', 'el-icon-sort-down', 256, '', 0, 0, 24, 1618637916, 1627540650, 0);
INSERT INTO `system_menu` VALUES (40, 23, 2, '添加 Markdown', '添加 Markdown', '/blog/article/addMd', '', 'views/blog/article/addMd', 'el-icon-sort-down', 258, '', 0, 0, 24, 1619072569, 1627540542, 0);
INSERT INTO `system_menu` VALUES (41, 22, 1, '文章分类', '文章分类', '/blog/article/category', '/blog/article/category/list', 'layout/components/RouterViewBox', 'el-icon-reading', 253, '7', 0, 0, 24, 1619342457, 1627539496, 0);
INSERT INTO `system_menu` VALUES (42, 41, 3, '添加分类', '添加分类', '/blog/article/category/add', '', 'views/blog/article/category/add', 'el-icon-sort-down', 256, '7', 0, 0, 24, 1619342897, 1627540591, 0);
INSERT INTO `system_menu` VALUES (43, 41, 3, '分类列表', '分类列表', '/blog/article/category/list', '', 'views/blog/article/category/list', 'el-icon-sort-down', 255, '7', 0, 0, 24, 1619342979, 1627538310, 0);
INSERT INTO `system_menu` VALUES (44, 41, 3, '编辑分类', '编辑分类', '/blog/article/category/upd', '', 'views/blog/article/category/upd', 'el-icon-sort-down', 257, '7', 1, 0, 24, 1619343044, 1627540595, 0);
INSERT INTO `system_menu` VALUES (45, 23, 2, '编辑 Markdown', '编辑 Markdown', '/blog/article/updMd', '', 'views/blog/article/updMd', 'el-icon-sort-down', 260, '7', 1, 0, 24, 1619359784, 1627540561, 0);
INSERT INTO `system_menu` VALUES (46, 0, 0, '上传云', '上传云', '/upload', '/upload/image/list', 'layout/index.vue', 'el-icon-upload', 6, '7', 0, 0, 24, 1619578268, 1627541005, 0);
INSERT INTO `system_menu` VALUES (47, 46, 1, '图床管理', '图床管理', '/upload/image', '/upload/image/list', 'layout/components/RouterViewBox', 'el-icon-picture', 255, '7', 0, 0, 24, 1619578758, 1627538415, 0);
INSERT INTO `system_menu` VALUES (48, 47, 2, '图床列表', '图床列表', '/upload/image/list', '', 'views/upload/image/list', 'el-icon-sort-down', 255, '7', 0, 0, 24, 1619578826, 1628740686, 0);
INSERT INTO `system_menu` VALUES (49, 47, 2, '上传图片', '上传图片', '/upload/image/add', '', 'views/upload/image/add', 'el-icon-sort-down', 256, '7', 0, 0, 24, 1619581515, 1627540672, 0);

-- ----------------------------
-- Table structure for system_role
-- ----------------------------
DROP TABLE IF EXISTS `system_role`;
CREATE TABLE `system_role`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '系统角色ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色名称',
  `description` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色描述',
  `created_user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '数据创建人',
  `created_at` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updated_at` int(11) NOT NULL DEFAULT 0 COMMENT '修改时间',
  `deleted_at` int(11) NOT NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id`(`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统角色表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of system_role
-- ----------------------------
INSERT INTO `system_role` VALUES (7, '超级管理员', '操控整个系统与数据', 24, 1618797694, 1618797818, 0);
INSERT INTO `system_role` VALUES (8, '博客管理员', '博客的操作员', 24, 1618797778, 1618797778, 0);
INSERT INTO `system_role` VALUES (9, 'aaa', 'aaaa', 24, 1618797835, 1618797835, 1618797878);

-- ----------------------------
-- Table structure for system_route
-- ----------------------------
DROP TABLE IF EXISTS `system_route`;
CREATE TABLE `system_route`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '系统路由ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '路由名称',
  `uri` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '路由 url',
  `method` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '请求类型',
  `role_list` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '授权角色组 逗号分隔 如 1,3',
  `is_auth` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否开启认证 0-否 1-是',
  `created_user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '数据创建人',
  `created_at` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updated_at` int(11) NOT NULL DEFAULT 0 COMMENT '修改时间',
  `deleted_at` int(11) NOT NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统路由表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of system_route
-- ----------------------------

-- ----------------------------
-- Table structure for upload_image
-- ----------------------------
DROP TABLE IF EXISTS `upload_image`;
CREATE TABLE `upload_image`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '上传图片ID',
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'url路径',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '实际路径',
  `created_user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '数据创建人',
  `created_at` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updated_at` int(11) NOT NULL DEFAULT 0 COMMENT '修改时间',
  `deleted_at` int(11) NOT NULL DEFAULT 0 COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 206 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '上传图片记录表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of upload_image
-- ----------------------------
INSERT INTO `upload_image` VALUES (2, 'http://127.0.0.1:8080/public/static/images/1a475874aeda2b32e7ce476a32c5224a.png', 'runtime/upload/images/1a475874aeda2b32e7ce476a32c5224a.png', 37, 1629857930, 1629857930, 0);
INSERT INTO `upload_image` VALUES (3, 'http://127.0.0.1:8080/public/static/images/dbc3dd3fb00baf0099c842b0aa169a9e.png', 'runtime/upload/images/dbc3dd3fb00baf0099c842b0aa169a9e.png', 37, 1629857978, 1629857978, 0);
INSERT INTO `upload_image` VALUES (4, 'http://127.0.0.1:8080/public/static/images/515f2a331603d1ffd6d6e3f602243521.png', 'runtime/upload/images/515f2a331603d1ffd6d6e3f602243521.png', 37, 1629857985, 1629857985, 0);

-- ----------------------------
-- Table structure for user_info
-- ----------------------------
DROP TABLE IF EXISTS `user_info`;
CREATE TABLE `user_info`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `role_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '角色ID',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '密码',
  `created_user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '数据创建人',
  `created_at` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updated_at` int(11) NOT NULL DEFAULT 0 COMMENT '修改时间',
  `deleted_at` int(11) NOT NULL DEFAULT 0 COMMENT '删除时间',
  `state` tinyint(4) NOT NULL DEFAULT 1 COMMENT '状态 0-删除 1-正常',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `state`(`state`) USING BTREE COMMENT '状态判断',
  INDEX `name`(`name`) USING BTREE COMMENT '用户名'
) ENGINE = InnoDB AUTO_INCREMENT = 27 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户信息表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_info
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
