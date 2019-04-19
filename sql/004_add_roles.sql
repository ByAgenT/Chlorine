CREATE TABLE IF NOT EXISTS member_role
(
    id        SERIAL PRIMARY KEY,
    role_name VARCHAR(100) NOT NULL,
    is_admin  BOOLEAN      NOT NULL
);

INSERT INTO member_role (id, role_name, is_admin)
VALUES (0, 'Member', FALSE),
       (1, 'Admin', TRUE);

ALTER TABLE IF EXISTS member
    DROP COLUMN is_admin;

ALTER TABLE IF EXISTS member
    ADD COLUMN role INT NOT NULL DEFAULT 0 REFERENCES member_role (id);

