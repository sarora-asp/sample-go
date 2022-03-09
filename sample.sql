CREATE TABLE IF NOT EXISTS user (
    id serial PRIMARY KEY,
    name VARCHAR (50) NOT NULL,
    email VARCHAR (100) UNIQUE NOT NULL,
    password VARCHAR (50) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
)

CREATE TABLE IF NOT EXISTS todo (
    	id serial PRIMARY KEY,
    	title VARCHAR (50) NOT NULL,
    	completed BOOLEAN DEFAULT false,
    	desc VARCHAR (50),
        user_id INTEGER,
    	created_at TIMESTAMP NOT NULL,
    	updated_at TIMESTAMP NOT NULL,
		due_date TIMESTAMP,
		completed_at TIMESTAMP
        FOREIGN KEY (user_id) REFERENCES user (id)
	)