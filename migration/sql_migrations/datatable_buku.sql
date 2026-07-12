
-- +migrate Up
DROP TRIGGER IF EXISTS modified_kategori_update ON Kategori;
DROP TRIGGER IF EXISTS modified_buku_update ON Buku;
DROP TRIGGER IF EXISTS modified_users_update ON Users;
DROP TRIGGER IF EXISTS trigger_set_thickness_buku ON Buku;

CREATE TABLE IF NOT EXISTS Kategori(
	id SERIAL PRIMARY KEY,
	name VARCHAR(200) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	created_by VARCHAR(200),
	modified_at TIMESTAMP,
	modified_by VARCHAR(200)
);

CREATE TABLE IF NOT EXISTS Users(
	id SERIAL PRIMARY KEY,
	username VARCHAR(200) NOT NULL,
	password VARCHAR(200) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	created_by VARCHAR(200),
	modified_at TIMESTAMP,
	modified_by VARCHAR(200)
);

CREATE TABLE IF NOT EXISTS Buku (
	id SERIAL PRIMARY KEY,
	title VARCHAR(200) NOT NULL,
	category_id INT REFERENCES Kategori(id) ON DELETE RESTRICT,
	description VARCHAR(255) NOT NULL,
	image_url TEXT,
	release_year INT,
	price INT,
	total_page INT,
	thickness VARCHAR(30),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	created_by VARCHAR(100),
	modified_at TIMESTAMP,
	modified_by VARCHAR(200)
);

-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION modified_buku_timestamp()
RETURNS TRIGGER AS $$
BEGIN
	NEW.modified_at = CURRENT_TIMESTAMP;

	RETURN NEW;
END;
$$ language 'plpgsql';
-- +migrate StatementEnd

-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION modified_kategori_timestamp()
RETURNS TRIGGER AS $$
BEGIN
	NEW.modified_at = CURRENT_TIMESTAMP;

	RETURN NEW;
END;
$$ language 'plpgsql';
-- +migrate StatementEnd

-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION modified_users_timestamp()
RETURNS TRIGGER AS $$
BEGIN
	NEW.modified_at = CURRENT_TIMESTAMP;

	RETURN NEW;
END;
$$ language 'plpgsql';
-- +migrate StatementEnd

-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION set_thickness_buku()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.total_page IS NULL THEN
        NEW.thickness = 'tipis';
    ELSIF NEW.total_page >= 100 THEN
        NEW.thickness = 'tebal';
    ELSE
        NEW.thickness = 'tipis';
    END IF;
    
    RETURN NEW;
END;
$$ language 'plpgsql';
-- +migrate StatementEnd


CREATE TRIGGER modified_buku_update
	BEFORE UPDATE ON Buku
	FOR EACH ROW
	EXECUTE PROCEDURE modified_buku_timestamp();

CREATE TRIGGER modified_kategori_update
	BEFORE UPDATE ON Kategori
	FOR EACH ROW
	EXECUTE PROCEDURE modified_kategori_timestamp();

CREATE TRIGGER modified_users_update
	BEFORE UPDATE ON Users
	FOR EACH ROW
	EXECUTE PROCEDURE modified_users_timestamp();

CREATE TRIGGER trigger_set_thickness_buku
    BEFORE INSERT OR UPDATE ON Buku
    FOR EACH ROW
    EXECUTE PROCEDURE set_thickness_buku();