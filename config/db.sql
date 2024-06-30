CREATE DATABASE Library;
USE Library;

CREATE TABLE Users (
    userID INT AUTO_INCREMENT NOT NULL,
    username varchar(255) DEFAULT NULL,
    userType enum('client','admin') DEFAULT 'client',
    name varchar(255) DEFAULT NULL,
    password varchar(1024) DEFAULT NULL,
    isSuperAdmin TINYINT(1) DEFAULT 0,
    hasAdminRequest TINYINT(1) DEFAULT 0,
    PRIMARY KEY (userID)
);

CREATE TABLE Books (
    bookID INT AUTO_INCREMENT NOT NULL,
    bookname varchar(1024) DEFAULT NULL,
    author varchar(1024) DEFAULT NULL,
    quantity INT DEFAULT 0,
    PRIMARY KEY (bookID)
);

CREATE TABLE Requests (
    requestID INT AUTO_INCREMENT NOT NULL,
    userID INT,
    bookID INT,
    isAccepted TINYINT(1) DEFAULT 0,
    isBorrowed TINYINT(1) DEFAULT 0,
    PRIMARY KEY (requestID),
    FOREIGN KEY (userID) REFERENCES Users(userID),
    FOREIGN KEY (bookID) REFERENCES Books(bookID)
);
