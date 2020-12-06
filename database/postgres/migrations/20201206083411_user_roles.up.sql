CREATE TABLE account_roles (
  user_id INT NOT NULL,
  role_id INT NOT NULL,
  created_at TIMESTAMP,
  PRIMARY KEY (user_id, role_id),
  FOREIGN KEY (role_id)
      REFERENCES roles (id),
  FOREIGN KEY (user_id)
      REFERENCES users (id)
);