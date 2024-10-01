DROP TABLE IF EXISTS books;
CREATE TABLE books (
  id         INT AUTO_INCREMENT NOT NULL,
  title      VARCHAR(128) NOT NULL,
  author     VARCHAR(255) NOT NULL,
  image      TEXT NOT NULL,
  PRIMARY KEY (`id`)
);
