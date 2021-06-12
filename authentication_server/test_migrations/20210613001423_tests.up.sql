CREATE TABLE users(
    id bigserial unique not null,
    username varchar unique not null,
    hash_password varchar not null
)