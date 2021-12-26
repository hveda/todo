CREATE DATABASE IF NOT EXISTS challenge_2_db;
USE challenge_2_db;

CREATE TABLE IF NOT EXISTS users (
  id          INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  lastname   VARCHAR(256) NOT NULL,
  firstname  VARCHAR(256)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO users(id, lastname, firstname) VALUES (1, "Yamada", "Takashi");