CREATE DATABASE mrhc;
USE mrhc;

CREATE TABLE accounts (
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE users (
  id INT NOT NULL AUTO_INCREMENT,
  account_id INT NOT NULL,
  name VARCHAR(255) NOT NULL,
  PRIMARY KEY (id)
);

INSERT INTO accounts (id, name) VALUES (1, 'account1');
INSERT INTO users (account_id, name) VALUES (1, 'user1');
INSERT INTO users (account_id, name) VALUES (1, 'user2');

INSERT INTO accounts (id, name) VALUES (2, 'account2');
INSERT INTO users (account_id, name) VALUES (2, 'user3');
INSERT INTO users (account_id, name) VALUES (2, 'user4')
