drop table if exists users;

create table users (
    id text not null primary key,
    name text not null
);

create table todos (
    id text not null primary key,
    user_id text not null,
    text text not null,
    done bool not null default false,

    foreign key (user_id) references users (id)
);
