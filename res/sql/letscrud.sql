DROP DATABASE IF EXISTS letscrud;
CREATE DATABASE letscrud;
USE letscrud;

CREATE TABLE customer (
    id SERIAL PRIMARY KEY,
    cpf VARCHAR(11) UNIQUE NOT NULL,
    name VARCHAR(200) NOT NULL,
    birthDate DATE
);

INSERT INTO customer (cpf, name, birthDate) VALUES (
    "71307366406",
    "Felipe da Silva",
    "2001-01-01"
);
SELECT * FROM customer;
SELECT * FROM customer WHERE id = 1;
SELECT * FROM customer WHERE cpf = "71307366406";
SELECT * FROM customer WHERE name LIKE "%lipe%";
SELECT * FROM customer WHERE birthDate = "2001-01-01";
UPDATE customer 
    SET cpf = "38219579056", name = "Felipe Santos", birthDate = "2002-02-02" 
    WHERE id = 1;
DELETE FROM customer WHERE id = 1;
