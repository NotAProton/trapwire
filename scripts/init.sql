CREATE DATABASE IF NOT EXISTS madb;
USE madb;

-- Create the 'wires' table
CREATE TABLE IF NOT EXISTS wires (
    id INT NOT NULL AUTO_INCREMENT,
    wire VARCHAR(255),
    dest VARCHAR(255),
    PRIMARY KEY (id)
);

-- Create the 'logs' table
CREATE TABLE IF NOT EXISTS logs (
    id INT NOT NULL AUTO_INCREMENT,
    time VARCHAR(255),
    ip VARCHAR(255),
    ua VARCHAR(255),
    referrer VARCHAR(255),
    wire VARCHAR(255),
    dest VARCHAR(255),
    PRIMARY KEY (id)
);
