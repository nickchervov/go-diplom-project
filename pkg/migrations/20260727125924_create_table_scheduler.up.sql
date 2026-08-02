CREATE TABLE scheduler
(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date varchar(8) NOT NULL DEFAULT '',
    title varchar(256) NOT NULL DEFAULT '',
    comment text,
    repeat varchar(128)
);

CREATE INDEX idx_scheduler_date ON scheduler (date);