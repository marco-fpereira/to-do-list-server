CREATE DATABASE TODOLIST;
USE TODOLIST;

DROP TABLE IF EXISTS ACCOUNT;
DROP TABLE IF EXISTS TASK;


CREATE TABLE ACCOUNT(
    UserId                  VARCHAR(256)
    Username                VARCHAR(256) NOT NULL,
    Password                VARCHAR(512) NOT NULL,
    PRIMARY KEY (UserId)
);

CREATE TABLE TASK(
    TaskId                  VARCHAR(36) NOT NULL,
    TaskMessage             VARCHAR(512),
    CreatedAt               TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    IsTaskCompleted         BOOLEAN
    PRIMARY KEY (TaskId)
    FOREIGN KEY (UserId) REFERENCES ACCOUNT(UserId)
);
