-- SQLite
SELECT id, name, description, location dateTime, user_id
FROM events

ALTER TABLE events
ADD location TEXT NOT NULL;

UPDATE events 
SET location = 'hehe'
WHERE ID = 2;