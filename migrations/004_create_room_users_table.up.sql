CREATE TABLE room_users(
    room_id BIGINT,
    user_id BIGINT,
    status INT8,
    PRIMARY KEY (room_id,user_id),
    FOREIGN KEY (room_id) REFERENCES rooms(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (user_id)REFERENCES  users(id) ON DELETE CASCADE ON UPDATE  CASCADE
);