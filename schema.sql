drop table if exists users;
create table users (
  id serial not null,
  email varchar(255) not null,
  name varchar(255) not null,
  password varchar(255) not null,
  primary key (id)
);

drop table if exists products;
create table products (
  id serial not null,
  name varchar(255) not null,
  description varchar(255) not null,
  price float not null,
  created_by int not null,
  primary key (id),
  foreign key (created_by) references users (id)
);

drop table if exists roles;
create table roles (
  id serial not null,
  name varchar(255) not null,
  primary key (id)
);

drop table if exists user_roles;
create table user_roles (
  id serial not null,
  user_id int not null,
  role_id int not null,
  primary key (id),
  foreign key (user_id) references users (id),
  foreign key (role_id) references roles (id)
);
