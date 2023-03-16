CREATE TABLE IF NOT EXISTS users
(
    id         INTEGER PRIMARY KEY AUTO_INCREMENT,
    first_name VARCHAR(30) NOT NULL,
    last_name  VARCHAR(50) NOT NULL,
    age        INTEGER(120),
    gender     VARCHAR(10),
    interests  VARCHAR(200),
    city       VARCHAR(70),
    email      VARCHAR(50) NOT NULL UNIQUE,
    password   VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
