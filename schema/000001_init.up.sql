CREATE TABLE music
(
    id serial not null unique,
    name varchar(255) not null,
    performer varchar(255) not null,
    release_date varchar(255) not null,
    genre varchar(255) not null
);