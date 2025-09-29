DROP TABLE IF EXISTS users;
DROP TAMBLE IF EXISTS posts;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    user_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

INSERT INTO users (email, password) VALUES ('test@test.com', '$2a$08$K8L1r45/3Y2.1V1eFgGdCOYbwBxHjKlMnOpQrStUvWxYzA1B2C3D4');
INSERT INTO users (email, password) VALUES ('test2@test.com', '$2a$08$K8L1r45/3Y2.1V1eFgGdCOYbwBxHjKlMnOpQrStUvWxYzA1B2C3D4');
INSERT INTO users (email, password) VALUES ('test3@test.com', '$2a$08$K8L1r45/3Y2.1V1eFgGdCOYbwBxHjKlMnOpQrStUvWxYzA1B2C3D4');


