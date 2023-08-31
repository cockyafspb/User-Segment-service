create table if not exists users
(
    user_id serial not null primary key
);

create table if not exists segments
(
    segment_id   serial  not null primary key,
    segment_name varchar not null unique
);

create table if not exists user_segment_pairs
(
    user_segment_id serial not null primary key,
    user_id         serial not null,
    foreign key (user_id) references users (user_id),
    segment_id      serial not null,
    foreign key (segment_id) references segments (segment_id),
    unique (user_id, segment_id)
);