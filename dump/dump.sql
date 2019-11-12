CREATE DATABASE IF NOT EXISTS echo;

USE echo;

CREATE TABLE IF NOT EXISTS product (
    product_id varchar(255) PRIMARY KEY,
    product_name varchar(255) NOT NULL,
    product_desc TEXT NOT NULL,
    product_image text NOT NULL,
    product_price DECIMAL(13, 2) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS content (
    content_id int PRIMARY KEY AUTO_INCREMENT,
    content_name varchar(255) NOT NULL,
    content_desc text NOT NULL,
    content_tags text NOT NULL,
    content_image text NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);