create schema authorization nameyourtime;

create extension "uuid-ossp";

--- TEST DB ---
create user test with password 'pass';
create database test_nameyourtime owner test;
grant nameyourtime to test;
alter user test with superuser;

create table nameyourtime.user (
                            id                varchar(40)  not null default uuid_generate_v4(),
                            name              varchar(128) not null,
                            email             varchar(128) not null,
                            password          varchar(255) not null,
                            status            varchar(50)  not null,
                            refresh_token     varchar(40)  not null default '',
                            refresh_token_exp timestamp    not null default now(),
                            created           timestamp    not null default now(),

                            primary key (id)
);
create unique index user_email_idx on nameyourtime.user(email);

create table nameyourtime.verification_code(
                                        user_id varchar(40) not null,
                                        code varchar(40) not null,
                                        code_exp timestamp not null,

                                        primary key (user_id)
);
create unique index auth_code_idx on nameyourtime.verification_code(code);