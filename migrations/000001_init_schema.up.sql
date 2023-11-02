create table users(
                      id int generated always as identity primary key,
                      name varchar(50) not null,
                      email varchar(50) not null unique,
                      hash_password text not null
);