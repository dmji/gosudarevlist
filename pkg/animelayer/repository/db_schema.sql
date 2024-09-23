CREATE TABLE items (
    guid TEXT PRIMARY KEY,
    title TEXT NOT NULL,
    category INT REFERENCES table_category(key) ON DELETE RESTRICT
);

CREATE TABLE table_category (
    key INT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    isexplicit BOOLEAN DEFAULT FALSE
);