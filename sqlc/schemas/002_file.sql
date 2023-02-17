CREATE TABLE IF NOT EXISTS file(
    file_uuid VARCHAR(256) PRIMARY KEY,
    title VARCHAR(256),
    passwdhash VARCHAR(512),
    access_token VARCHAR(40) NOT NULL UNIQUE,
    encrypted BOOLEAN NOT NULL DEFAULT FALSE,
    file_size INTEGER,
    encryption_iv BYTEA DEFAULT NULL,
    upload_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_seen TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    viewcount INTEGER NOT NULL
);