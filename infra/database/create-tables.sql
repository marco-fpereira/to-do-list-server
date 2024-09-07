CREATE DATABASE IF NOT EXISTS TODOLIST;
USE TODOLIST;

DROP TABLE IF EXISTS ACCOUNT;
DROP TABLE IF EXISTS TASK;

CREATE TABLE ACCOUNT(
    UserId                  VARCHAR(36) NOT NULL,
    Username                VARCHAR(256) NOT NULL,
    Password                VARCHAR(512) NOT NULL,
    PRIMARY KEY (UserId)
);

CREATE TABLE TASK(
    TaskId                  VARCHAR(36) NOT NULL,
    TaskMessage             VARCHAR(512),
    CreatedAt               TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    IsTaskCompleted         BOOLEAN,
    UserId                  VARCHAR(36) NOT NULL,
    PRIMARY KEY (TaskId),
    FOREIGN KEY (UserId) REFERENCES ACCOUNT(UserId)
);
