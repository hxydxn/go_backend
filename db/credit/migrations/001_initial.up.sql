CREATE TABLE IF NOT EXISTS credits (
    id SERIAL PRIMARY KEY,
    ssrn INT NOT NULL UNIQUE,
    project_id TEXT NOT NULL,
    std TEXT NOT NULL,
    methodology TEXT NOT NULL,
    region TEXT NOT NULL,
    storage_method TEXT NOT NULL,
    method TEXT NOT NULL,
    emission_type TEXT NOT NULL,
    catergory TEXT NOT NULL,
    uri TEXT NOT NULL,
    beneficiary TEXT NOT NULL
);

-- TODO: Switch from TEXT to Var Char (TEXT is shit and inefficient)