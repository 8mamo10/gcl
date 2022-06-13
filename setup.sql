CREATE DATABASE IF NOT EXISTS mrhc;
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
  email VARCHAR(255) NOT NULL,
  PRIMARY KEY (id)
);

INSERT INTO accounts (id, name) VALUES (1, 'account1');
INSERT INTO users (account_id, name, email) VALUES (1, 'user1', 'user1@8mamo10.net');
INSERT INTO users (account_id, name, email) VALUES (1, 'user2', 'user2@8mamo10.net');

INSERT INTO accounts (id, name) VALUES (2, 'account2');
INSERT INTO users (account_id, name, email) VALUES (2, 'user3', 'user3@8mamo10.net');
INSERT INTO users (account_id, name, email) VALUES (2, 'user4', 'user4@8mamo10.net')
