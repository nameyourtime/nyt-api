create schema authorization nameyourtime;

create extension "uuid-ossp";

--- TEST DB ---
create user test with password 'pass';
create database test_nameyourtime owner test;
grant nameyourtime to test;
alter user test with superuser;