CREATE TABLE `account`
(
    `id`         BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `created_at` DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` DATETIME        NULL COMMENT '删除时间',
    `account_id` VARCHAR(128)    NOT NULL COMMENT '唯一账户ID',
    `username`   VARCHAR(64)     NOT NULL COMMENT '用户帐号',
    `email`      VARCHAR(128)    NULL COMMENT '邮箱',
    `password`   VARCHAR(256)    NOT NULL COMMENT '密码md5',
    `salt`       VARCHAR(256)    NOT NULL COMMENT '盐',
    `status`     VARCHAR(32)     NOT NULL COMMENT '帐号状态',
    PRIMARY KEY (`id`),
    UNIQUE INDEX `uniq_account_id` (`account_id`),
    UNIQUE INDEX `uniq_username` (`username`, `deleted_at`),
    UNIQUE INDEX `uniq_email` (`email`, `deleted_at`)
) ENGINE = INNODB
  DEFAULT CHARSET = utf8mb4
    COMMENT '用户帐号表';

CREATE TABLE `model`
(
    `id`           BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `created_at`   DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`   DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`   DATETIME        NULL COMMENT '删除时间',
    `model_id`     VARCHAR(128)    NOT NULL COMMENT '唯一模型id',
    `platform`     VARCHAR(64)     NOT NULL COMMENT '模型所属平台',
    `name`         VARCHAR(64)     NOT NULL COMMENT '模型名称',
    `display_name` VARCHAR(64)     NOT NULL COMMENT '展示的名字',
    PRIMARY KEY (`id`),
    UNIQUE INDEX `uniq_model_id` (`model_id`)
) ENGINE = INNODB
  DEFAULT CHARSET = utf8mb4
    COMMENT '模型表';

CREATE TABLE `account_model`
(
    `id`          BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `created_at`  DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`  DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`  DATETIME        NULL COMMENT '删除时间',
    `relation_id` VARCHAR(128)    NOT NULL COMMENT '唯一关系ID',
    `account_id`  VARCHAR(128)    NOT NULL COMMENT '唯一账户ID',
    `model_id`    VARCHAR(128)    NOT NULL COMMENT '唯一模型id',
    PRIMARY KEY (`id`),
    UNIQUE INDEX `uniq_rid` (`relation_id`),
    UNIQUE INDEX `uniq_acc_model` (`account_id`, `model_id`, `deleted_at`)
) ENGINE = INNODB
  DEFAULT CHARSET = utf8mb4
    COMMENT '账户-模型关联表';