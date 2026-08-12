ALTER TABLE guests
    ADD COLUMN IF NOT EXISTS confirmed BOOLEAN;

UPDATE guests SET
    confirmed = CASE
        WHEN responded THEN attending
        ELSE NULL
    END;

ALTER TABLE guests
    DROP COLUMN IF EXISTS responded,
    DROP COLUMN IF EXISTS attending;

ALTER TABLE plus_ones
    DROP COLUMN IF EXISTS attending;
