CREATE DATABASE IF NOT EXISTS `user-test` CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
create table users (id int, name varchar(255), email VARCHAR(50));
insert into users values (1, 'Yamada', 'test@example.com');