CREATE TABLE courses (
	id serial PRIMARY KEY,
    user_id INT NOT NULL, -- teacher id
    FOREIGN KEY (user_id)
      REFERENCES users (id),
	name VARCHAR ( 50 ) NOT NULL, -- course name
	identitiy VARCHAR ( 50 ) UNIQUE NOT NULL, -- identitiy of course
    valence SMALLINT DEFAULT 25 NOT NULL CHECK (valence > 0), -- Capacity for each course, default is 25
    time VARCHAR ( 150 ) NOT NULL, -- course time and class detail 
    created_at TIMESTAMP NOT NULL
);
CREATE INDEX course_identitiy ON courses(identitiy);