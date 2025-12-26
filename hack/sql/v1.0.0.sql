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