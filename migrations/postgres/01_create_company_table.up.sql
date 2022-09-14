CREATE TABLE IF NOT EXISTS books (
    "id" uuid PRIMARY KEY,
    "name" VARCHAR(255) UNIQUE NOT NULL,
    "about" TEXT NOT NULL,
    "isbn" VARCHAR(255) NOT NULL
);
