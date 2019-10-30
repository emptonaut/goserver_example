-- quick and dirty database schema init
-- column names should be reflected in user.go and session.go

CREATE TABLE users(
    id integer PRIMARY KEY AUTOINCREMENT,
    username text UNIQUE NOT NULL,
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
