CREATE TABLE IF NOT EXISTS users
(
    Id                     TEXT PRIMARY KEY,
    Email                  TEXT UNIQUE NOT NULL,
    Password               TEXT        NOT NULL,
    AccessToken            TEXT,
    RefreshToken           TEXT,
    Verified               BOOLEAN   DEFAULT FALSE,
    LastResetSentAt        TEXT,
    LastVerificationSentAt TEXT,
    CreatedAt              TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt              TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS categories
(
    Id        TEXT PRIMARY KEY,
    Name      TEXT NOT NULL,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS profiles
(
    Id        TEXT PRIMARY KEY,
    Name      TEXT,
    Avatar    TEXT,
    UserId    TEXT UNIQUE NOT NULL,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_profiles_users FOREIGN KEY (UserId) REFERENCES users (Id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS farms
(
    Id         TEXT PRIMARY KEY,
    Name       TEXT        NOT NULL,
    CategoryId TEXT UNIQUE NOT NULL,
    UserId     TEXT UNIQUE NOT NULL,
    CreatedAt  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_farms_users FOREIGN KEY (UserId) REFERENCES users (Id),
    CONSTRAINT fk_farms_categories FOREIGN KEY (CategoryId) REFERENCES categories (Id)
)