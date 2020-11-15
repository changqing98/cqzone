CREATE TABLE `user`
(
    id        INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键自增',
    user_id   INT UNSIGNED NOT NULL COMMENT '用户id',
    nickname  VARCHAR(40)  NOT NULL COMMENT '昵称',
    mobile    VARCHAR(11)  NOT NULL DEFAULT '' COMMENT '手机号',
    email     VARCHAR(30)  NOT NULL COMMENT '邮箱',
    password  VARCHAR(60)  NOT NULL COMMENT '密码',
    created_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY uni_user_mobile (`mobile`) USING BTREE,
    UNIQUE KEY uni_user_email (`email`) USING BTREE,
    UNIQUE KEY uni_user_user_id (`user_id`) USING BTREE
) DEFAULT CHARSET = utf8mb4;
