CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    email      VARCHAR(100) UNIQUE NOT NULL,
    password   VARCHAR(100)        NOT NULL,
    username   VARCHAR(255)        NOT NULL,
    first_name VARCHAR(100)        NOT NULL,
    last_name  VARCHAR(100)        NOT NULL,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE profiles
(
    user_id             INT PRIMARY KEY,
    city                VARCHAR(255),
    birth_date          DATE,
    relationship_status VARCHAR(255),
    status_text         VARCHAR(255),
    profile_emoji       TEXT,


    created_at          TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE chats
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(255),
    type       VARCHAR(20) NOT NULL CHECK ( type IN ('personal', 'group') ),
    creator_id INT,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (creator_id) REFERENCES users (id) ON DELETE SET NULL
);

CREATE TYPE chat_roles AS ENUM('admin', 'owner', 'moderator', 'user');

CREATE TABLE chat_members
(
    chat_id    INT        NOT NULL,
    user_id    INT        NOT NULL,
    role       chat_roles NOT NULL DEFAULT 'user',

    created_at TIMESTAMPTZ         DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ         DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (chat_id) REFERENCES chats (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,

    PRIMARY KEY (chat_id, user_id)
);

CREATE TABLE messages
(
    id                  SERIAL PRIMARY KEY,
    chat_id             INT NOT NULL,
    sender_id           INT,
    reply_to_message_id INT,
    content             TEXT,

    created_at          TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at          TIMESTAMPTZ,

    FOREIGN KEY (sender_id) REFERENCES users (id) ON DELETE SET NULL,
    FOREIGN KEY (chat_id) REFERENCES chats (id) ON DELETE CASCADE,
    FOREIGN KEY (reply_to_message_id) REFERENCES messages (id) ON DELETE SET NULL
);

CREATE TABLE message_files
(
    id         SERIAL PRIMARY KEY,
    message_id INT          NOT NULL,
    name       VARCHAR(255) NOT NULL,
    path       VARCHAR(255) NOT NULL,
    size       BIGINT       NOT NULL,

    FOREIGN KEY (message_id) REFERENCES messages (id) ON DELETE CASCADE,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE reaction_types
(
    id    SERIAL PRIMARY KEY,
    code  VARCHAR(50) NOT NULL UNIQUE,
    emoji TEXT        NOT NULL
);

CREATE TABLE message_reactions
(
    message_id  INT NOT NULL,
    user_id     INT NOT NULL,
    reaction_id INT NOT NULL,
    created_at  TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (message_id) REFERENCES messages (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (reaction_id) REFERENCES reaction_types (id) ON DELETE RESTRICT,
    PRIMARY KEY (message_id, user_id, reaction_id)
);
