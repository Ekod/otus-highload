CREATE TABLE IF NOT EXISTS posts(
  id INTEGER PRIMARY KEY AUTO_INCREMENT,
  user_id INTEGER NOT NULL,
  content VARCHAR(500),
  FOREIGN KEY (user_id) REFERENCES users(id)
);