CREATE TABLE room_users(
    room_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL ,
    PRIMARY KEY (room_id,user_id),
    FOREIGN KEY (room_id) REFERENCES rooms(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (user_id)REFERENCES  users(id) ON DELETE CASCADE ON UPDATE  CASCADE
);