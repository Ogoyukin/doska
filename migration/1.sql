CREATE DATABASE doska CHARACTER SET UTF8mb4 COLLATE utf8mb4_unicode_ci;

create table items
(
    id int auto_increment primary key,
    price int,
    title varchar(255),
    description varchar(255),
    img varchar(255)
);

