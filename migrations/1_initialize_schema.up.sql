create table "user" (
    id bigserial primary key,
    email varchar(255) unique not null,
    password varchar(255) not null,
    name varchar(255),
    surname varchar(255),
    phone varchar(32),
    city varchar(255),
    avatar_url text
);
