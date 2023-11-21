create table if not exists users
(
    id         bigint unsigned auto_increment primary key,
    first_name varchar(256) not null,
    last_name  varchar(256) not null,
    phone      varchar(256) not null,
    birth_date datetime     not null
    ) ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4;

