create table `gomall_user`(
    `id`            bigint unsigned not null                            ,
    `name`          varchar(255)    not null        default '' comment '用户名',
    `gender`        tinyint(3)      not null        default 0 comment '用户性别',
    `mobile`        varchar(255)    not null        default '' comment '电话号码',
    `password`      varchar(255)    not null        default '' comment '用户密码',
    `create_time`   timestamp       not null        comment '创建时间',
    `updated_time`  timestamp       not null        comment '修改时间',
    primary key (`id`)
)ENGINE = InnoDB
Default CHARSET = utf8mb4;