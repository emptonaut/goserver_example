-- quick and dirty database schema init
-- column names should be reflected in Users.go

CREATE TABLE users(
    id integer PRIMARY KEY AUTOINCREMENT,
    username text NOT NULL,
    password text NOT NULL,
    salt text NOT NULL
);

CREATE TABLE sessions (
    id integer PRIMARY KEY AUTOINCREMENT,
    userid integer NOT NULL,
    token text NOT NULL,
    origin text NOT NULL,
    expires text NOT NULL
);
