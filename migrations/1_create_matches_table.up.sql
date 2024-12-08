CREATE TABLE matches(
    id TEXT PRIMARY KEY,
    date TIMESTAMP WITH TIME ZONE,
    name VARCHAR(255),
    team_a VARCHAR(255),
    team_b VARCHAR(255),
    banner_url VARCHAR(255),
    sanko_tx BIGINT UNIQUE,
    arbitrum_tx BIGINT UNIQUE,
    ethereum_tx BIGINT UNIQUE,
    base_tx BIGINT UNIQUE,
    ape_tx BIGINT UNIQUE,
    creation_date TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    update_date TIMESTAMP WITH TIME ZONE
);
