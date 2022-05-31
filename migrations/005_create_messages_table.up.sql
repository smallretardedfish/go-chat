CREATE TABLE messages(
  id BIGSERIAL PRIMARY KEY,
  text TEXT NOT NULL,
  owner_id BIGINT,
  room_id BIGINT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY(owner_id) REFERENCES users(id)  ON DELETE SET NULL,
  FOREIGN KEY(room_id) REFERENCES rooms(id) ON DELETE CASCADE
);