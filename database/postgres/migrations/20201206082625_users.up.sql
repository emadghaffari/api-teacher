CREATE TABLE users (
	id serial PRIMARY KEY,
	name VARCHAR ( 50 ) NOT NULL, -- name of user
	lname VARCHAR ( 50 ) NOT NULL, -- last name(family) of user
	identitiy VARCHAR ( 50 ) UNIQUE NOT NULL, -- identitiy of user
	password VARCHAR ( 50 ) NOT NULL -- password
);

CREATE INDEX user_identitiy ON users(identitiy);