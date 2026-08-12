-- Seed de teste: convidados pendentes para validar o fluxo de confirmação

INSERT INTO guests (name, responded, attending) VALUES
    ('João Pedro Almeida', false, false),
    ('Fernanda Costa Lima', false, false),
    ('Ricardo Souza Martins', false, false),
    ('Camila Ribeiro Santos', false, false),
    ('Bruno Henrique Oliveira', false, false),
    ('Patrícia Mendes Carvalho', false, false),
    ('Eduardo Tavares Nunes', false, false),
    ('Juliana Ferreira Rocha', false, false),
    ('Marcelo Antunes Pereira', false, false),
    ('Beatriz Cardoso Lopes', false, false);

-- Acompanhantes (plus_ones) vinculados a alguns dos convidados acima
INSERT INTO plus_ones (guest_id, name, attending)
SELECT id, 'Cônjuge de ' || name, false
FROM guests
WHERE name IN ('João Pedro Almeida', 'Fernanda Costa Lima', 'Ricardo Souza Martins');

INSERT INTO plus_ones (guest_id, name, attending)
SELECT id, 'Filho(a) de ' || name, false
FROM guests
WHERE name IN ('João Pedro Almeida', 'Camila Ribeiro Santos');
