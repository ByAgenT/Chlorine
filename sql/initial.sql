CREATE TABLE IF NOT EXISTS room_config
(
  id               SERIAL PRIMARY KEY,
  songs_per_member INT,
  max_members      INT       NOT NULL DEFAULT 20,
  created_date     TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS room
(
  id            SERIAL PRIMARY KEY,
  spotify_token VARCHAR(255) NOT NULL,
  config_id     INT          NOT NULL REFERENCES room_config (id),
  created_date  TIMESTAMP    NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS member
(
  id           SERIAL PRIMARY KEY,
  name         VARCHAR(255) NOT NULL,
  session      VARCHAR(255) NOT NULL,
  room_id      INT          NOT NULL REFERENCES room (id),
  is_admin     BIT(1)       NOT NULL,
  created_date TIMESTAMP    NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS song
(
  id              SERIAL PRIMARY KEY,
  spotify_id      VARCHAR(255) NOT NULL,
  room_id         INT          NOT NULL REFERENCES room (id),
  prev_song_id    INT REFERENCES song (id),
  next_song_id    INT REFERENCES song (id),
  member_added_id INT          NOT NULL,
  created_date    TIMESTAMP    NOT NULL DEFAULT NOW()
);