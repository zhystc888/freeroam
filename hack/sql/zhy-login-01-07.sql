INSERT INTO `free_system_config` (`config_code`, `config_value`, `config_desc`)
VALUES
    ('auth.session.idle_timeout', '1800', '登录空闲过期时间（秒）'),
    ('auth.session.max_lifetime', '86400', '登录最大过期时间'),
    ('auth.allow_multi_login', 'false', '是否允许多点登录');

-- 枚举类型（只保留必要字段）
INSERT INTO `free_system_enum_type` (`enum_type`, `enum_type_desc`)
VALUES ('member_status', '成员表-状态');

-- 枚举值数据
INSERT INTO `free_system_enum_data`
(`enum_type`, `enum_code`, `enum_value`, `enum_label`, `enum_value_desc`, `sort`)
VALUES
    ('member_status', 'enable',    '1', '禁用',  NULL, 0),
    ('member_status', 'disable',   '2', '启用',  NULL, 0),
    ('member_status', 'resigned',  '3', '离职',  NULL, 0);