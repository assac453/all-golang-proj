package model

var EXEC = `
	CREATE TABLE task(
		id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		title VARCHAR(255) NOT NULL,
		description TEXT,
		completed BOOL
	);
`
