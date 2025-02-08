CREATE TABLE users
(
    id            INT AUTO_INCREMENT PRIMARY KEY COMMENT '用户ID',
    email         VARCHAR(255) UNIQUE NOT NULL COMMENT '用户邮箱',
    phone         VARCHAR(20) UNIQUE COMMENT '用户手机号（可选）',
    username      VARCHAR(50)         NOT NULL COMMENT '用户名',
    password_hash VARCHAR(255)        NOT NULL COMMENT '密码的哈希值',
    avatar_url    VARCHAR(255)        NOT NULL COMMENT '用户头像URL',
    bio           TEXT COMMENT '用户简介（可选）',
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '注册时间',
    updated_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
    last_login_at TIMESTAMP COMMENT '上次登录时间',
    status        TINYINT   DEFAULT 1 COMMENT '账户状态，1 表示正常，0 表示禁用',
    deleted_at    TIMESTAMP           NULL COMMENT '软删除字段，非 NULL 表示已删除'
) COMMENT ='用户表，存储平台用户的基本信息';


CREATE TABLE roles
(
    id          INT AUTO_INCREMENT PRIMARY KEY COMMENT '角色ID',
    name        VARCHAR(50) UNIQUE NOT NULL COMMENT '角色名称，例如 "user", "admin"',
    description TEXT COMMENT '角色描述',
    deleted_at  TIMESTAMP          NULL COMMENT '软删除字段，非 NULL 表示已删除'
) COMMENT ='角色表，存储平台中的角色信息';

CREATE TABLE permissions
(
    id          INT AUTO_INCREMENT PRIMARY KEY COMMENT '权限ID',
    name        VARCHAR(50) UNIQUE NOT NULL COMMENT '权限名称，例如 "create_link", "view_stats"',
    description TEXT COMMENT '权限描述',
    deleted_at  TIMESTAMP          NULL COMMENT '软删除字段，非 NULL 表示该权限已被删除'
) COMMENT ='权限表，存储平台的权限信息';

CREATE TABLE user_roles
(
    user_id     INT       NOT NULL COMMENT '用户ID',
    role_id     INT       NOT NULL COMMENT '角色ID',
    assigned_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '角色分配时间',
    deleted_at  TIMESTAMP NULL COMMENT '软删除字段，非 NULL 表示该角色与用户的关联已删除',
    PRIMARY KEY (user_id, role_id)
) COMMENT ='用户与角色的关系表，记录每个用户拥有哪些角色';

CREATE TABLE role_permissions
(
    role_id       INT       NOT NULL COMMENT '角色ID',
    permission_id INT       NOT NULL COMMENT '权限ID',
    granted_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '权限授予时间',
    deleted_at    TIMESTAMP NULL COMMENT '软删除字段，非 NULL 表示该权限与角色的关联已删除',
    PRIMARY KEY (role_id, permission_id)
) COMMENT ='角色与权限的关系表，记录每个角色拥有的权限';

CREATE TABLE user_permissions
(
    user_id       INT       NOT NULL COMMENT '用户ID',
    permission_id INT       NOT NULL COMMENT '权限ID',
    granted_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '权限授予时间',
    deleted_at    TIMESTAMP NULL COMMENT '软删除字段，非 NULL 表示该权限与用户的关联已删除',
    PRIMARY KEY (user_id, permission_id)
) COMMENT ='用户权限表，记录每个用户拥有的权限';

CREATE TABLE access_logs
(
    id            INT AUTO_INCREMENT PRIMARY KEY COMMENT '日志ID',
    user_id       INT       NOT NULL COMMENT '用户ID',
    permission_id INT       NOT NULL COMMENT '权限ID',
    action        TINYINT NOT NULL DEFAULT 1  COMMENT '表示权限是被授予还是被撤销1-granted 2-revoked"',
    action_time   TIMESTAMP                   DEFAULT CURRENT_TIMESTAMP COMMENT '操作时间',
    deleted_at    TIMESTAMP NULL COMMENT '软删除字段，非 NULL 表示该日志已删除'
) COMMENT ='访问日志表，记录用户操作权限的日志';
