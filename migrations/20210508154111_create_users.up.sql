CREATE TABLE users (
    id bigserial not null primary key,
    name varchar not null,
    email varchar not null unique,
    encrypted_password varchar,

--Not main fields
    sex varchar,
    room_id bigserial,
    faculty_id bigserial
);