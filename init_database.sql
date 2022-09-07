CREATE TABLE users
(
    Id                     TEXT PRIMARY KEY,
    Email                  TEXT UNIQUE NOT NULL,
    Password               TEXT        NOT NULL,
    Token                  TEXT,
    LastResetSentAt        TEXT,
    LastVerificationSentAt TEXT,
    Verified               BOOLEAN DEFAULT FALSE,
    CreatedAt              TEXT,
    UpdatedAt              TEXT
);