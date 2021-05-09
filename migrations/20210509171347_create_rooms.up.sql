CREATE TABLE rooms (
    id bigserial not null primary key,
    capacity bigserial not null,
    free_capacity bigserial not null,
    hostel_id bigserial not null,
    room_sex varchar not null
);