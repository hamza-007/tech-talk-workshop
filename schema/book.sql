CREATE TABLE books (
    id VARCHAR(255) PRIMARY KEY UNIQUE,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    year int NOT NULL,
    edition VARCHAR(255) NOT NULL,
    price decimal(2,3) NOT NULL    
);