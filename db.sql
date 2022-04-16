CREATE SCHEMA pastebin;

USE pastebin;

CREATE TABLE contents (
  shortlink VARCHAR(60) PRIMARY KEY,
  content TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  expiry_in_minutes INTEGER NOT NULL DEFAULT 0
);

