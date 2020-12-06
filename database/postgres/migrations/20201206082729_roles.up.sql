CREATE TABLE roles(
   id serial PRIMARY KEY, -- role id
   name VARCHAR (255) UNIQUE NOT NULL -- role name like teacher,student,...
);