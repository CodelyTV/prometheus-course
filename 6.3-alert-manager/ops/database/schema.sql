CREATE USER 'mysqld_exporter'@'localhost' IDENTIFIED BY 'password' WITH MAX_USER_CONNECTIONS 3;

CREATE TABLE IF NOT EXISTS `users` (
     `id` INT NOT NULL AUTO_INCREMENT,
     `name` VARCHAR(255),
     `email` VARCHAR(255),
     PRIMARY KEY (`id`)
);