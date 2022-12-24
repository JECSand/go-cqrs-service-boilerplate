CREATE EXTENSION IF NOT EXISTS citext;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS groups CASCADE;
DROP TABLE IF EXISTS group_memberships CASCADE;
DROP TABLE IF EXISTS blacklists CASCADE;


CREATE TABLE users
(
    id              UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    username        VARCHAR(250)  NOT NULL CHECK ( username <> '' ),
    email           VARCHAR(250)  NOT NULL CHECK ( email <> '' ),
    password        VARCHAR(250) NOT NULL CHECK ( password <> '' ),
    root            BOOLEAN       NOT NULL,
    active          BOOLEAN       NOT NULL,
    created_at      TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE groups
(
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name        VARCHAR(250)  NOT NULL CHECK ( name <> '' ),
    description VARCHAR(250) NOT NULL CHECK ( description <> '' ),
    creator_id  UUID NOT NULL,
    active       BOOLEAN       NOT NULL,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (creator_id) REFERENCES users(id)
);

CREATE TABLE group_memberships
(
    id  UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    name        VARCHAR(250)  NOT NULL CHECK ( name <> '' ),
    description VARCHAR(250) NOT NULL CHECK ( description <> '' ),
    user_id     UUID NOT NULL,
    group_id    UUID NOT NULL,
    status      BOOLEAN       NOT NULL,
    role        BOOLEAN       NOT NULL,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (group_id) REFERENCES groups(id)
);

CREATE TABLE blacklists
(
    id  UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    token        VARCHAR(2500)  NOT NULL CHECK ( token <> '' ),
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);