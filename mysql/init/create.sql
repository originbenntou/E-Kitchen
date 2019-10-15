DROP DATABASE IF EXISTS e_kitchen;
CREATE DATABASE e_kitchen;

use e_kitchen

DROP TABLE IF EXISTS users;
CREATE TABLE users (
  id INT unsigned NOT NULL auto_increment,
  email VARCHAR(255) NOT NULL,
  password VARCHAR(1023) NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);
