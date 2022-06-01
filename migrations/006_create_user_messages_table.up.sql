CREATE TABLE user_messages (
   user_id BIGINT NOT NULL,
   message_id BIGINT NOT NULL,
   status INT8 NOT NULL,
   PRIMARY KEY (user_id,message_id),
   FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL ON UPDATE CASCADE ,
   FOREIGN KEY (message_id) REFERENCES messages(id) ON DELETE CASCADE ON UPDATE CASCADE
);