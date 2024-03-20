CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    surname varchar(255) not null,
    age int not null,
    password_hash varchar(255) not null,
    email varchar(255) not null,
    is_admin boolean not null default false
) 