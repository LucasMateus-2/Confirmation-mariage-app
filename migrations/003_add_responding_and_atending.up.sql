ALTER TABLE guests
    ADD COLUMN IF NOT EXISTS responded BOOLEAN NOT NULL DEFAULT false,
    ADD COLUMN IF NOT EXISTS attending BOOLEAN NOT NULL DEFAULT false;

-- Migra dados existentes da coluna antiga "confirmed" (tri-state), se ela existir
DO $$
BEGIN
    IF EXISTS (
        SELECT 1 FROM information_schema.columns
        WHERE table_name = 'guests' AND column_name = 'confirmed'
    ) THEN
        UPDATE guests SET
            responded = (confirmed IS NOT NULL),
            attending = COALESCE(confirmed, false);

        ALTER TABLE guests DROP COLUMN confirmed;
    END IF;
END $$;

ALTER TABLE plus_ones
    ADD COLUMN IF NOT EXISTS attending BOOLEAN NOT NULL DEFAULT false;
