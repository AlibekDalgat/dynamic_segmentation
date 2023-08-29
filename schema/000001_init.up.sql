CREATE TABLE segments
(
    id serial not null unique,
    name varchar(255) not null primary key
);

CREATE TABLE users_in_segments
(
    user_id integer not null,
    segment_name varchar(255) references segments(name) on delete cascade not null
);