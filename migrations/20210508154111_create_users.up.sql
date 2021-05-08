CREATE TABLE users (
    id bigserial not null primary key,
    name varchar not null,
    email varchar not null,
    room_num bigserial,
    encrypted_password varchar
);