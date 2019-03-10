CREATE TABLE IF NOT EXISTS spotify_token
(
  id            SERIAL PRIMARY KEY,
  access_token  VARCHAR(512) NOT NULL,
  expiry        TIMESTAMP    NOT NULL,
  refresh_token VARCHAR(512) NOT NULL,
  token_type    VARCHAR(255) NOT NULL
);

ALTER TABLE IF EXISTS room
  DROP COLUMN spotify_token;
ALTER TABLE IF EXISTS room
  ADD COLUMN spotify_token INT NULL REFERENCES spotify_token (id);