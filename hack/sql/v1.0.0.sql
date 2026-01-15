--- 枚举表
CREATE TABLE `free_system_enum_type` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `enum_type` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '枚举类型，如 user_type',
  `enum_type_desc` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '枚举类型说明',
  `is_enabled` tinyint DEFAULT '1' COMMENT '是否启用0:否,1:是',
  `is_deleted` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否删除0:否,1:是',
  `create_by` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建人',
  `update_by` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改人',
  `delete_by` bigint unsigned NOT NULL DEFAULT '0' COMMENT '删除人',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_enum` (`enum_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='枚举类型';

CREATE TABLE `free_system_enum_data` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `enum_type` varchar(64) COLLATE utf8mb4_general_ci NOT NULL COMMENT '枚举类型，如 user_type',
  `enum_code` varchar(64) COLLATE utf8mb4_general_ci NOT NULL COMMENT '枚举编码，如 ADMIN',
  `enum_value` varchar(64) COLLATE utf8mb4_general_ci NOT NULL COMMENT '枚举值，如 1',
  `enum_label` varchar(128) COLLATE utf8mb4_general_ci NOT NULL COMMENT '前端展示文本',
  `enum_value_desc` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '枚举值说明',
  `sort` int DEFAULT '0' COMMENT '顺序',
  `is_enabled` tinyint DEFAULT '1' COMMENT '是否启用0:否,1:是',
  `is_deleted` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否删除0:否,1:是',
  `create_by` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建人',
  `update_by` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改人',
  `delete_by` bigint unsigned NOT NULL DEFAULT '0' COMMENT '删除人',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_enum` (`enum_type`,`enum_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='枚举数据';

-- 系统配置
CREATE TABLE `free_system_config` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `config_code` varchar(64) COLLATE utf8mb4_general_ci NOT NULL COMMENT '配置编码，如 token_validity_period',
  `config_value` varchar(64) COLLATE utf8mb4_general_ci NOT NULL COMMENT '配置值，如 60',
  `config_desc` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '配置说明，如 token有效时长(分钟)',
  `is_deleted` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否删除0:否,1:是',
  `create_by` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建人',
  `update_by` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改人',
  `delete_by` bigint unsigned NOT NULL DEFAULT '0' COMMENT '删除人',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_enum` (`config_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='系统配置';

-- 角色表
CREATE TABLE `free_role` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `code` varchar(64) NOT NULL COMMENT '角色编码(唯一)',
  `name` varchar(64) NOT NULL COMMENT '角色名称',
  `status` varchar(64) NOT NULL COMMENT '状态:role_status',
  `is_system` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否系统内置0:否,1:是',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `is_deleted` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否删除0:否,1:是',
  `create_by` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建人',
  `update_by` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改人',
  `delete_by` bigint unsigned NOT NULL DEFAULT '0' COMMENT '删除人',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='角色';

-- 组织表
CREATE TABLE `free_org` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `parent_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '父组织ID, 0为顶级',
  `name` varchar(30) NOT NULL COMMENT '组织名称(支持模糊搜索)',
  `full_name` varchar(128) NOT NULL DEFAULT '' COMMENT '组织全称',
  `code` varchar(20) NOT NULL COMMENT '组织编码(唯一标识, 精确搜索)',
  `category` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '分类:1部门2公司',
  `status` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '状态:1启用0禁用',
  `sort` int NOT NULL DEFAULT '0' COMMENT '同级排序(拖拽排序用)',
  `path` varchar(512) NOT NULL DEFAULT '/' COMMENT '物化路径, 如 /1/10/ (用于子树查询与防环)',
  `is_deleted` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '数据状态0正常1删除',
  `create_by` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建人',
  `update_by` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改人',
  `delete_by` bigint unsigned NOT NULL DEFAULT '0' COMMENT '删除人',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_parent_sort_del` (`parent_id`, `sort`, `is_deleted`),
  KEY `idx_category_del` (`category`, `is_deleted`),
  KEY `idx_status_del` (`status`, `is_deleted`),
  KEY `idx_path_del` (`path`, `is_deleted`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='组织';

-- 职务表(含数据权限配置)
CREATE TABLE `free_position` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(20) NOT NULL COMMENT '职务名称(<=20)',
  `status` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '状态:1启用0禁用',
  `data_scope` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '数据权限(仅配置):1本人2本部门3本部门及子部门4全体',
  `is_deleted` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '数据状态0正常1删除',
  `create_by` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建人',
  `update_by` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改人',
  `delete_by` bigint unsigned NOT NULL DEFAULT '0' COMMENT '删除人',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_name_del` (`name`, `is_deleted`),
  KEY `idx_status_del` (`status`, `is_deleted`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='职务(含数据权限配置)';

-- 职务-组织关联表(职务关联组织)
CREATE TABLE `free_position_org` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `position_id` bigint unsigned NOT NULL COMMENT '职务ID',
  `org_id` bigint unsigned NOT NULL COMMENT '组织ID',
  `is_deleted` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '数据状态0正常1删除',
  `create_by` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建人',
  `update_by` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改人',
  `delete_by` bigint unsigned NOT NULL DEFAULT '0' COMMENT '删除人',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_pos_org_del` (`position_id`, `org_id`, `is_deleted`),
  KEY `idx_org_del` (`org_id`, `is_deleted`),
  KEY `idx_pos_del` (`position_id`, `is_deleted`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='职务-组织关联(职务关联组织)';

-- 职务-角色关联表
CREATE TABLE `free_position_role` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `position_id` bigint unsigned NOT NULL COMMENT '职务ID',
  `role_id` bigint unsigned NOT NULL COMMENT '角色ID',
  `is_deleted` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '数据状态0正常1删除',
  `create_by` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建人',
  `update_by` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改人',
  `delete_by` bigint unsigned NOT NULL DEFAULT '0' COMMENT '删除人',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_pos_role_del` (`position_id`, `role_id`, `is_deleted`),
  KEY `idx_pos_del` (`position_id`, `is_deleted`),
  KEY `idx_role_del` (`role_id`, `is_deleted`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='职务-角色关联';