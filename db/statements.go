package db

const user_table = `CREATE TABLE IF NOT EXISTS users(
id VARCHAR PRIMARY KEY, 
name VARCHAR, email VARCHAR, 
phone INTEGER, 
location VARCHAR)`