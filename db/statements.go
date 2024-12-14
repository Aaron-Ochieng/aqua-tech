package db

const user_table = `
	CREATE TABLE
    IF NOT EXISTS users (
        id VARCHAR(255) PRIMARY KEY,
        name VARCHAR(255),
        email VARCHAR(255) UNIQUE,
        phone INTEGER,
        location VARCHAR(255),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
`

const device_table = `
	CREATE TABLE
    IF NOT EXISTS device (
        id VARCHAR(255) PRIMARY KEY UNIQUE,
        pond_id VARCHAR(255) UNIQUE,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        owner_id VARCHAR(255),
        FOREIGN KEY (owner_id) REFERENCES USERS (id)
    );
`
