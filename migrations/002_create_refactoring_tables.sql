

CREATE TABLE IF NOT EXISTS guests (
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(150) NOT NULL,
    responded  BOOLEAN DEFAULT FALSE,
    attending  BOOLEAN DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS plus_ones (
    id        SERIAL PRIMARY KEY,
    guest_id  INT NOT NULL REFERENCES guests(id) ON DELETE CASCADE,
    name      VARCHAR(150) NOT NULL,
    attending BOOLEAN DEFAULT FALSE
);
