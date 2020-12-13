CREATE TABLE courses (
	id serial PRIMARY KEY,
    user_id INT NOT NULL, -- teacher id
    FOREIGN KEY (user_id)
      REFERENCES users (id),
	name VARCHAR ( 50 ) NOT NULL, -- course name
	identitiy VARCHAR ( 50 ) UNIQUE NOT NULL, -- identitiy of course
    valence SMALLINT DEFAULT 25 NOT NULL CHECK (valence >= 0), -- Capacity for each course, default is 25
    value SMALLINT DEFAULT 1 NOT NULL CHECK (value >= 1), -- Capacity for each course, default is 25
    start_at TIME NOT NULL, -- course time and class detail 
    end_at TIME NOT NULL, -- course time and class detail 
    description VARCHAR ( 150 ) NOT NULL, -- course time and class detail 
    created_at TIMESTAMP NOT NULL
);
CREATE INDEX course_identitiy ON courses(identitiy);