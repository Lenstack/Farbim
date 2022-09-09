CREATE TABLE IF NOT EXISTS users
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

CREATE TABLE IF NOT EXISTS profiles
(
    Id        TEXT PRIMARY KEY,
    Name      TEXT,
    Avatar    TEXT,
    UserId    TEXT UNIQUE NOT NULL,
    CreatedAt TEXT,
    UpdatedAt TEXT,
    CONSTRAINT fk_profiles_users FOREIGN KEY (UserId) REFERENCES users (Id)
)
