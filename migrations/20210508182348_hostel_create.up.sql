CREATE TABLE hostels (
    id bigserial not null primary key,
    description varchar not null unique,
    faculty_id bigserial not null
);