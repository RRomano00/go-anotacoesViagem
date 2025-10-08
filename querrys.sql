ALTER TABLE travel
ADD COLUMN creation_at TIMESTAMPTZ DEFAULT NOW();

CREATE TABLE notes (
    id SERIAL PRIMARY KEY,
    content TEXT NOT NULL,
    travel_id INTEGER NOT NULL,
    
    -- Colunas de auditoria que você já conhece
    creation_at TIMESTAMPTZ DEFAULT NOW(),
    update_date TIMESTAMPTZ DEFAULT NOW(),

    -- A linha mágica que cria a relação
    CONSTRAINT note_pkey
        FOREIGN KEY(note_travel_id_fkey) 
        REFERENCES travel(id)
        ON DELETE CASCADE
);