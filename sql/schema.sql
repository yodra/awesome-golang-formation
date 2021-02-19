CREATE TABLE movies
(
    id     int(8) unsigned NOT NULL AUTO_INCREMENT,
    name   VARCHAR(255) NOT NULL,
    year   VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,

    PRIMARY KEY (id)

) CHARACTER SET utf8mb4
  COLLATE utf8mb4_bin;
