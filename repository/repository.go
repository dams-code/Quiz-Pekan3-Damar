package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"quiz-pekan3-damar/structbuku"
)

func TambahUser(db *sql.DB, setUser structbuku.Users) (HasilTambahUser structbuku.Users, err error) {
	queryInsert := `
        INSERT INTO users (username, password, created_by)
        VALUES ($1, $2, $1)
        RETURNING id, username, created_at;
    `

	if db == nil {
		return HasilTambahUser, fmt.Errorf("koneksi database nil")
	}

	err = db.QueryRow(queryInsert, setUser.Username, setUser.Password).Scan(
		&HasilTambahUser.ID,
		&HasilTambahUser.Username,
		&HasilTambahUser.CreatedAt,
	)

	if err != nil {
		return HasilTambahUser, err
	}

	return HasilTambahUser, nil
}

func GetUsername(db *sql.DB, Username string) (HasilGetUser structbuku.Users, err error) {
	querySelect := `
        SELECT id, username, password 
        FROM Users
        WHERE username = $1;
    `

	err = db.QueryRow(querySelect, Username).Scan(
		&HasilGetUser.ID,
		&HasilGetUser.Username,
		&HasilGetUser.Password,
	)

	if err != nil {
		return structbuku.Users{}, err
	}

	return HasilGetUser, nil
}

func GetKategori(db *sql.DB) (HasilAllKategori []structbuku.Kategori, err error) {
	querySelect := `
		SELECT 	id, name, created_at, 
				created_by, modified_at, modified_by
		FROM Kategori
	`

	rows, err := db.Query(querySelect)

	if err != nil {

		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var setKategori structbuku.Kategori

		err = rows.Scan(
			&setKategori.ID, &setKategori.Name,
			&setKategori.CreatedAt, &setKategori.CreatedBy,
			&setKategori.ModifiedAt, &setKategori.ModifiedAt,
		)

		if err != nil {
			return nil, err
		}

		HasilAllKategori = append(HasilAllKategori, setKategori)
	}

	if HasilAllKategori == nil {
		HasilAllKategori = []structbuku.Kategori{}
	}

	return HasilAllKategori, nil
}

func GetKategoriID(db *sql.DB, IdKategori int) (HasilGetKategori structbuku.Kategori, err error) {
	querySelect := `
		SELECT 	id, name, created_at, 
				created_by, modified_at, modified_by
		FROM Kategori
		WHERE id = $1
	`

	err = db.QueryRow(querySelect, IdKategori).Scan(
		&HasilGetKategori.ID, &HasilGetKategori.Name, &HasilGetKategori.CreatedAt,
		&HasilGetKategori.CreatedBy, &HasilGetKategori.ModifiedAt, &HasilGetKategori.ModifiedBy,
	)

	if err != nil {
		return structbuku.Kategori{}, err
	}

	return HasilGetKategori, nil
}

func UpdateKategori(db *sql.DB, IdKategori int, setKategori structbuku.Kategori) (HasilUpdateKategori structbuku.Kategori, err error) {
	queryUpdate := `
		UPDATE Kategori
        SET name = $1, modified_by = $2, modified_at = NOW()
        WHERE id = $3
        RETURNING id, name, created_at, created_by, modified_at, modified_by
	`

	err = db.QueryRow(queryUpdate,
		setKategori.Name,
		setKategori.ModifiedBy, IdKategori,
	).Scan(
		&HasilUpdateKategori.ID, &HasilUpdateKategori.Name, &HasilUpdateKategori.CreatedAt, &HasilUpdateKategori.CreatedBy,
		&HasilUpdateKategori.ModifiedAt, &HasilUpdateKategori.ModifiedBy,
	)

	if err != nil {
		if err != sql.ErrNoRows {
			return structbuku.Kategori{}, fmt.Errorf("Id-%d Kategori Tidak ditemukan", IdKategori)
		}
		return structbuku.Kategori{}, err
	}

	return HasilUpdateKategori, nil
}

func TambahKategori(db *sql.DB, setKategori structbuku.Kategori) (HasilTambahKategori structbuku.Kategori, err error) {
	queryInsert := `
		INSERT INTO Kategori (name, created_by)
		VALUES($1, $2)
		RETURNING id, name, created_at, created_by
	`

	err = db.QueryRow(queryInsert, setKategori.Name, setKategori.CreatedBy).Scan(&HasilTambahKategori.ID, &HasilTambahKategori.Name, &HasilTambahKategori.CreatedAt, &HasilTambahKategori.CreatedBy)

	if err != nil {
		return structbuku.Kategori{}, err
	}

	return HasilTambahKategori, nil

}

func HapusKategori(db *sql.DB, IdKategori int) (err error) {
	queryDelete := `
		DELETE FROM Kategori 
		WHERE id = $1
		RETURNING id, name
	`

	HasilDelete, err := db.Exec(queryDelete, IdKategori)

	rowsAffected, err := HasilDelete.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("kategori ID %d tidak ada di database", IdKategori)
	}

	return nil
}

func GetBuku(db *sql.DB) (HasilGetBuku []structbuku.Buku, err error) {
	querySelect := `
		SELECT 	b.id, b.title, b.description, b.category_id, k.name AS categoryname,
				b.image_url, b.release_year, b.price, b.total_page,
				b.thickness, b.created_at, b.created_by, 
				b.modified_at, b.modified_by
		FROM Buku as b
		LEFT JOIN Kategori k on b.category_id = k.id
	`

	rows, err := db.Query(querySelect)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Daftar buku kosong / tidak ada")
		}
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var setBuku structbuku.Buku

		err := rows.Scan(
			&setBuku.ID, &setBuku.Title, &setBuku.Description, &setBuku.CategoryID, &setBuku.CategoryName,
			&setBuku.ImageURL, &setBuku.ReleaseYear, &setBuku.Price, &setBuku.TotalPage,
			&setBuku.Thickness, &setBuku.CreatedAt, &setBuku.CreatedBy, &setBuku.ModifiedAt, &setBuku.ModifiedBy,
		)

		if err != nil {
			return nil, err
		}

		HasilGetBuku = append(HasilGetBuku, setBuku)
	}

	if HasilGetBuku == nil {
		HasilGetBuku = []structbuku.Buku{}
	}

	return HasilGetBuku, nil
}

func GetBukuBerdasarkanKategori(db *sql.DB, IdKategoriBuku int) (HasilGetBukuBerdasarkanKategori []structbuku.Buku, err error) {
	queryKategoriBuku := `
		SELECT 
			b.id, b.title, b.category_id, k.name AS category_name, 
			b.description, b.image_url, b.release_year, b.price, 
			b.total_page, b.thickness
		FROM Buku b
		LEFT JOIN kategori k ON b.category_id = k.id
		WHERE b.category_id = $1;
	`

	rows, err := db.Query(queryKategoriBuku, IdKategoriBuku)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Id-%d Kategori Buku Tidak ditemukan", IdKategoriBuku)
		}
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var setKategoriBuku structbuku.Buku

		err = rows.Scan(
			&setKategoriBuku.ID, &setKategoriBuku.Title,
			&setKategoriBuku.CategoryID, &setKategoriBuku.CategoryName,
			&setKategoriBuku.Description, &setKategoriBuku.ImageURL,
			&setKategoriBuku.ReleaseYear, &setKategoriBuku.Price,
			&setKategoriBuku.TotalPage, &setKategoriBuku.Thickness,
		)

		if err != nil {
			return nil, err
		}

		HasilGetBukuBerdasarkanKategori = append(HasilGetBukuBerdasarkanKategori, setKategoriBuku)
	}

	if HasilGetBukuBerdasarkanKategori == nil {
		HasilGetBukuBerdasarkanKategori = []structbuku.Buku{}
	}

	return HasilGetBukuBerdasarkanKategori, nil

}

func GetBukuID(db *sql.DB, IdBuku int) (HasilGetBuku structbuku.Buku, err error) {
	querySelect := `
		SELECT 	b.id, b.title, b.description, b.category_id, k.name AS categoryname,
				b.image_url, b.release_year, b.price, b.total_page,
				b.thickness, b.created_at, b.created_by, 
				b.modified_at, b.modified_by
		FROM Buku as b
		LEFT JOIN Kategori k on b.category_id = k.id
		WHERE b.id = $1
	`

	err = db.QueryRow(querySelect, IdBuku).Scan(
		&HasilGetBuku.ID, &HasilGetBuku.Title, &HasilGetBuku.Description,
		&HasilGetBuku.CategoryID, &HasilGetBuku.CategoryName, &HasilGetBuku.ImageURL,
		&HasilGetBuku.ReleaseYear, &HasilGetBuku.Price, &HasilGetBuku.TotalPage,
		&HasilGetBuku.Thickness, &HasilGetBuku.CreatedAt, &HasilGetBuku.CreatedBy,
		&HasilGetBuku.ModifiedAt, &HasilGetBuku.ModifiedBy,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return structbuku.Buku{}, fmt.Errorf("Id-%d Buku Tidak Ada", IdBuku)
		}
		return structbuku.Buku{}, err
	}

	return HasilGetBuku, nil
}

func TambahBuku(db *sql.DB, setBuku structbuku.Buku) (HasilTambahBuku structbuku.Buku, err error) {
	queryInsert := `
		INSERT INTO Buku (title, category_id, description, image_url, 
					release_year, price, total_page, thickness, created_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, title, category_id, description, image_url, release_year, price, total_page, thickness, created_at, created_by, modified_at, modified_by;
	`

	err = db.QueryRow(queryInsert, setBuku.Title, setBuku.CategoryID,
		setBuku.Description, setBuku.ImageURL,
		setBuku.ReleaseYear, setBuku.Price,
		setBuku.TotalPage, setBuku.Thickness,
		setBuku.CreatedBy).Scan(&HasilTambahBuku.ID, &HasilTambahBuku.Title, &HasilTambahBuku.CategoryID,
		&HasilTambahBuku.Description, &HasilTambahBuku.ImageURL, &HasilTambahBuku.ReleaseYear, &HasilTambahBuku.Price,
		&HasilTambahBuku.TotalPage, &HasilTambahBuku.Thickness, &HasilTambahBuku.CreatedAt, &HasilTambahBuku.CreatedBy,
		&HasilTambahBuku.ModifiedAt, &HasilTambahBuku.ModifiedBy,
	)

	if err != nil {
		return structbuku.Buku{}, err
	}

	return HasilTambahBuku, nil
}

func UpdateBuku(db *sql.DB, IdBuku int, setBuku structbuku.Buku) (HasilUpdateBuku structbuku.Buku, err error) {
	queryUpdate := `
		UPDATE Buku
		SET
			title = $1, category_id = $2, description = $3, 
			image_url = $4, release_year = $5, price = $6, 
			total_page = $7, thickness = $8, modified_by = $9,
			modified_at = NOW()
		WHERE id = $10
		RETURNING 	id, title, category_id, description, image_url, release_year, price, 
					total_page, thickness, created_at, created_by, modified_at, modified_by
	`

	err = db.QueryRow(queryUpdate,
		setBuku.Title, setBuku.CategoryID, setBuku.Description,
		setBuku.ImageURL, setBuku.ReleaseYear, setBuku.Price,
		setBuku.TotalPage, setBuku.Thickness, setBuku.ModifiedBy, IdBuku,
	).Scan(
		&HasilUpdateBuku.ID, &HasilUpdateBuku.Title, &HasilUpdateBuku.CategoryID,
		&HasilUpdateBuku.Description, &HasilUpdateBuku.ImageURL, &HasilUpdateBuku.ReleaseYear,
		&HasilUpdateBuku.Price, &HasilUpdateBuku.TotalPage, &HasilUpdateBuku.Thickness,
		&HasilUpdateBuku.CreatedAt, &HasilUpdateBuku.CreatedBy, &HasilUpdateBuku.ModifiedAt, &HasilUpdateBuku.ModifiedBy,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return structbuku.Buku{}, fmt.Errorf("Id-%d Buku Tidak ditemukan", IdBuku)
		}
		return structbuku.Buku{}, err
	}

	return HasilUpdateBuku, nil
}

func HapusBuku(db *sql.DB, IdBuku int) (err error) {
	queryDelete := `
		DELETE FROM Buku
		WHERE id = $1
		RETURNING id, title
	`
	HasilHapusBuku, err := db.Exec(queryDelete, IdBuku)
	if err != nil {
		return err
	}

	rowsAffected, err := HasilHapusBuku.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("Id-%d Buku Tidak Ada", IdBuku)
	}

	return nil
}
