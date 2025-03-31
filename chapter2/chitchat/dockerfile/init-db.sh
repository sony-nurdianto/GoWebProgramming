#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    GRANT ALL PRIVILEGES ON DATABASE $POSTGRES_DB TO $POSTGRES_USER;

    CREATE TABLE users (
        id         SERIAL PRIMARY KEY,
        uuid       VARCHAR(64) NOT NULL UNIQUE,
        name       VARCHAR(255),
        email      VARCHAR(255) NOT NULL UNIQUE,
        password   VARCHAR(255) NOT NULL,
        created_at TIMESTAMP NOT NULL
    );

    CREATE TABLE sessions (
        id         SERIAL PRIMARY KEY,
        uuid       VARCHAR(64) NOT NULL UNIQUE,
        email      VARCHAR(255),
        user_id    INTEGER REFERENCES users(id),
        created_at TIMESTAMP NOT NULL
    );

    CREATE TABLE threads (
        id         SERIAL PRIMARY KEY,
        uuid       VARCHAR(64) NOT NULL UNIQUE,
        topic      TEXT,
        user_id    INTEGER REFERENCES users(id),
        created_at TIMESTAMP NOT NULL
    );

    CREATE TABLE posts (
        id         SERIAL PRIMARY KEY,
        uuid       VARCHAR(64) NOT NULL UNIQUE,
        body       TEXT,
        user_id    INTEGER REFERENCES users(id),
        thread_id  INTEGER REFERENCES threads(id),
        created_at TIMESTAMP NOT NULL
    );
EOSQL
