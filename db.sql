CREATE SCHEMA pastebin;

USE pastebin;

CREATE TABLE pastes (
  shortlink VARCHAR(60) PRIMARY KEY,
  text TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  expiry_in_minutes INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE url_visit_histories (
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    shortlink VARCHAR(255) NOT NULL,
    address VARCHAR(50),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
