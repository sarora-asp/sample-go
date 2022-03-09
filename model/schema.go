package model

const Schema = `
	CREATE TABLE IF NOT EXISTS users (
		id serial PRIMARY KEY,
		name VARCHAR (50) NOT NULL,
		email VARCHAR (100) UNIQUE NOT NULL,
		password VARCHAR (50) NOT NULL,
		created_at TIMESTAMPTZ DEFAULT Now(),
		updated_at TIMESTAMPTZ DEFAULT Now()
	);
`
