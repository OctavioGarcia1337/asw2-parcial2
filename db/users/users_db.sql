-- MySQL Script generated by MySQL Workbench
-- Mon Dec 12 22:48:13 2022
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema users_db
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema users_db
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `users_db` DEFAULT CHARACTER SET utf8 ;
USE `users_db` ;

-- -----------------------------------------------------
-- Table `users_db`.`users`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `users_db`.`users` ;

CREATE TABLE IF NOT EXISTS `users_db`.`users` (
      `id` INT NOT NULL AUTO_INCREMENT,
      `username` VARCHAR(45) NOT NULL,
    `password` VARCHAR(45) NOT NULL,
    `email` VARCHAR(255) NOT NULL,
    `first_name` VARCHAR(255) NOT NULL,
    `last_name` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE INDEX `username_UNIQUE` (`username` ASC) VISIBLE)
    ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `users_db`.`messages`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `users_db`.`messages` ;

CREATE TABLE IF NOT EXISTS `users_db`.`messages` (
     `id` INT NOT NULL AUTO_INCREMENT,
     `body` VARCHAR(500) NOT NULL,
    `user_id` INT NOT NULL,
    `item_id` VARCHAR(500) NULL,
    `created_at` DATETIME NOT NULL,
    `system` BOOL NOT NULL,
    PRIMARY KEY (`id`),
    INDEX `fk_messages_users_idx` (`user_id` ASC) VISIBLE,
    CONSTRAINT `fk_messages_users`
    FOREIGN KEY (`user_id`)
    REFERENCES `users_db`.`users` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION)
    ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;


INSERT INTO `users_db`.`users`(username, password, email, first_name, last_name) VALUES("system", "_DEFAULT", "_SYSTEM", "System", "Admin"); 
