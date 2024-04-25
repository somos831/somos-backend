-- schema.sql
-- defines the schema for the `somos_events` database

CREATE TABLE IF NOT EXISTS categories (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS events (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(50) NOT NULL,
    description VARCHAR(1000),
    category_id INT NOT NULL,
    location VARCHAR(200),
    PRIMARY KEY (id),
    FOREIGN KEY (category_id) REFERENCES categories(id)
);
