CREATE TABLE IF NOT EXISTS authors (
	id INTEGER PRIMARY KEY,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
	name TEXT NOT NULL,
	bio TEXT
);

