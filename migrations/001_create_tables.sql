-- migrations/001_create_tables.sql

CREATE TABLE IF NOT EXISTS users (
    id       SERIAL PRIMARY KEY,
    email    VARCHAR(150) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS guests (
    id        SERIAL PRIMARY KEY,
    name      VARCHAR(150) NOT NULL,
    confirmed BOOLEAN DEFAULT NULL  -- NULL = pendente, TRUE = confirmado, FALSE = recusou
);

CREATE TABLE IF NOT EXISTS plus_ones (
    id       SERIAL PRIMARY KEY,
    guest_id INT NOT NULL REFERENCES guests(id) ON DELETE CASCADE,
    name     VARCHAR(150) NOT NULL
);

-- Seed: inserir os noivos como usuários admin
-- Rode isso manualmente após criar as tabelas, substituindo os hashes reais
-- INSERT INTO users (name, email, password) VALUES
--   ('Lucas', 'lucas@email.com', '$2a$10$...hash_bcrypt...'),
--   ('Noiva', 'noiva@email.com', '$2a$10$...hash_bcrypt...');
INSERT INTO users (email, password) VALUES
  ('lucas.venanciorossi@outlook.com', '$2a$10$CapD/kuiW6NER9PUYJk8DOc2KCIV1f86vJchYh49HUZBBmEgxepjS'),
  ('mariaclara@gmail.com', '$2a$10$.q1qunUqoly8HZD1gW/ge.saNzUs.TXWf4k1RW.law18sPHf0arnK');
