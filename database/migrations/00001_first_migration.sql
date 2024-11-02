-- +goose NO TRANSACTION
-- +goose Up
-- CREATE DATABASE  bookstore;

CREATE TABLE if not exists books (
    isbn VARCHAR PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    "edition" VARCHAR(255) NOT NULL,
    publisher VARCHAR(255) NOT NULL,
    source VARCHAR(255) NOT NULL,
    cost DECIMAL(10, 2) NOT NULL
);

CREATE TABLE if not exists  users (
    id UUID PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    "password" VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    firstname VARCHAR(255) NOT NULL,
    lastname VARCHAR(255),
    "role" VARCHAR(100) NOT NULL DEFAULT 'ROLE_USER',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE if not exists  "copy" (
    -- adding foreign key for isbn
    book_isbn VARCHAR(255) NOT NULL,
    user_id UUID NOT NULL,
    barcode CHAR(255) PRIMARY KEY NOT NULL,
    available BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (book_isbn) REFERENCES books(isbn)
);

CREATE TABLE   if not exists "borrow" (
    -- adding foreign key for isbn
    copy_barcode char(255) NOT NULL,
    user_id UUID NOT NULL,
    borrowed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    returned_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (copy_barcode) REFERENCES "copy"(barcode)
);


-- +goose NO TRANSACTION
-- +goose Down
-- +goose StatementBegin

DROP TABLE borrow;
DROP TABLE "copy";
DROP TABLE "users";
DROP TABLE books;

-- DROP DATABASE bookstore;
-- +goose StatementEnd
