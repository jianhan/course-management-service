-- +migrate Up

-- MySQL Script generated by MySQL Workbench
-- Sat 03 Feb 2018 21:09:10 AEST
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL,ALLOW_INVALID_DATES';

-- -----------------------------------------------------
-- Schema course-management-service
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema course-management-service
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `course-management-service` DEFAULT CHARACTER SET utf8 ;
USE `course-management-service` ;

-- -----------------------------------------------------
-- Table `course-management-service`.`courses`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `course-management-service`.`courses` (
  `id` CHAR(36) NOT NULL,
  `name` VARCHAR(512) NOT NULL,
  `slug` VARCHAR(512) NOT NULL,
  `display_order` INT NOT NULL DEFAULT 0,
  `description` TEXT NOT NULL,
  `visible` TINYINT NOT NULL DEFAULT 1,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP(),
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC),
  UNIQUE INDEX `name_UNIQUE` (`name` ASC),
  UNIQUE INDEX `slug_UNIQUE` (`slug` ASC),
  FULLTEXT INDEX `description_FULLTEXT` (`description` ASC))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_bin;


-- -----------------------------------------------------
-- Table `course-management-service`.`categories`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `course-management-service`.`categories` (
  `id` CHAR(36) NOT NULL,
  `name` VARCHAR(512) NOT NULL,
  `description` VARCHAR(1024) NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP(),
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `idtable1_UNIQUE` (`id` ASC),
  UNIQUE INDEX `name_UNIQUE` (`name` ASC),
  FULLTEXT INDEX `description_FULLTEXT` (`description` ASC))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_bin;


-- -----------------------------------------------------
-- Table `course-management-service`.`courses_categories`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `course-management-service`.`courses_categories` (
  `courses_id` CHAR(36) NOT NULL,
  `categories_id` CHAR(36) NOT NULL,
  `display_order` INT NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT NOW(),
  PRIMARY KEY (`courses_id`, `categories_id`),
  INDEX `fk_courses_has_categories_categories1_idx` (`categories_id` ASC),
  INDEX `fk_courses_has_categories_courses1_idx` (`courses_id` ASC),
  CONSTRAINT `fk_courses_has_categories_courses1`
    FOREIGN KEY (`courses_id`)
    REFERENCES `course-management-service`.`courses` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_courses_has_categories_categories1`
    FOREIGN KEY (`categories_id`)
    REFERENCES `course-management-service`.`categories` (`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_bin;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;

-- +migrate Down
