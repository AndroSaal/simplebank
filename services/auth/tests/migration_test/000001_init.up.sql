CREATE TABLE IF NOT EXISTS migration_tests (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    passsword_hash VARCHAR(255) NOT NULL
);

INSERT INTO migration_tests (id, email, passsword_hash)  
VALUES (1, 'test', 'test-secret') 
ON CONFLICT DO NOTHING;