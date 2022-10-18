CREATE TABLE IF NOT EXISTS _permissions
(
    Id        TEXT PRIMARY KEY,
    Service   TEXT,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS _roles
(
    Id            TEXT PRIMARY KEY,
    Name          TEXT,
    PermissionsId TEXT,
    CreatedAt     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt     TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS _users
(
    Id              TEXT PRIMARY KEY,
    Email           TEXT UNIQUE NOT NULL,
    Password        TEXT        NOT NULL,
    TokenKey        TEXT,
    Verified        BOOLEAN   DEFAULT FALSE,
    RolesId         TEXT,
    Code            TEXT,
    LastResetSentAt TEXT,
    CreatedAt       TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt       TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS _profiles
(
    Id        TEXT PRIMARY KEY,
    Name      TEXT,
    Avatar    TEXT,
    UserId    TEXT UNIQUE NOT NULL,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_profiles_users FOREIGN KEY (UserId) REFERENCES _users (Id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS _sessions
(
    Id           TEXT PRIMARY KEY,
    Type         TEXT,
    Token        TEXT,
    Blocked      BOOLEAN   DEFAULT FALSE,
    ClientOrigin TEXT,
    ExpiredAt    TEXT,
    UserId       TEXT,
    CreatedAt    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt    TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS _settings
(
    Id        TEXT PRIMARY KEY,
    Key       TEXT,
    Value     JSON,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);