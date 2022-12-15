CREATE EXTENSION IF NOT EXISTS citext;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
DROP TABLE IF EXISTS products CASCADE;


CREATE TABLE users
(
    id              UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    username        VARCHAR(250)  NOT NULL CHECK ( username <> '' ),
    email           VARCHAR(250)  NOT NULL CHECK ( email <> '' ),
    password        VARCHAR(250) NOT NULL CHECK ( password <> '' ),
    root            NUMERIC       NOT NULL,
    active          NUMERIC       NOT NULL,
    created_at      TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE groups
(
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name        VARCHAR(250)  NOT NULL CHECK ( name <> '' ),
    description VARCHAR(250) NOT NULL CHECK ( description <> '' ),
    creator_id  UUID FOREIGN KEY NOT NULL,
    active       NUMERIC       NOT NULL,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE group_memberships
(
    id  UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    name        VARCHAR(250)  NOT NULL CHECK ( name <> '' ),
    description VARCHAR(250) NOT NULL CHECK ( description <> '' ),
    user_id     UUID FOREIGN KEY NOT NULL,
    group_id    UUID FOREIGN KEY NOT NULL,
    status      NUMERIC       NOT NULL,
    role        NUMERIC       NOT NULL,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE blacklists
(
    id  UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    token        VARCHAR(2500)  NOT NULL CHECK ( token <> '' ),
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
);