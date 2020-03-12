DROP DATABASE IF EXISTS resource;
CREATE DATABASE resource;

use resource

DROP TABLE IF EXISTS shop;
CREATE TABLE shop (
  id INT unsigned NOT NULL auto_increment,
  name VARCHAR(255) NOT NULL DEFAULT "unknown",
  status INT unsigned NOT NULL DEFAULT 0,
  url VARCHAR(255) NOT NULL DEFAULT "",
  user_id INT unsigned NOT NULL DEFAULT 0,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

DROP TABLE IF EXISTS tag;
CREATE TABLE tag (
  id INT unsigned NOT NULL auto_increment,
  name VARCHAR(255) NOT NULL DEFAULT "unknown",
  status INT unsigned NOT NULL DEFAULT 0,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

DROP TABLE IF EXISTS tag_map;
CREATE TABLE tag_map (
  id INT unsigned NOT NULL auto_increment,
  shop_id INT unsigned NOT NULL,
  tag_id INT unsigned NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  FOREIGN KEY (shop_id) REFERENCES shop(id) ON DELETE CASCADE,
  FOREIGN KEY (tag_id) REFERENCES tag(id) ON DELETE CASCADE
);
