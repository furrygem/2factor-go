CREATE TABLE users(
    id bigserial unique not null,
    username varchar unique not null,
    discord_username varchar unique not null,
    first_name varchar,
    hash_password varchar not null,   
)