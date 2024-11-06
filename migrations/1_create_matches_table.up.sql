CREATE TABLE matches(
    id TEXT PRIMARY KEY,
    date TIMESTAMP WITH TIME ZONE,
    name VARCHAR(255),
    team_a VARCHAR(255),
    team_b VARCHAR(255),
    banner_url VARCHAR(255)
);
