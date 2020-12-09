CREATE TABLE user_course (
  user_id INT NOT NULL,
  course_id INT NOT NULL,
  created_at TIMESTAMP NOT NULL,
  PRIMARY KEY (user_id, course_id),
  FOREIGN KEY (course_id)
      REFERENCES courses (id),
  FOREIGN KEY (user_id)
      REFERENCES users (id)
);