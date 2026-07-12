package structbuku

import "time"

type BukuReq struct {
	Title       string  `json:"title" binding:"required"`
	CategoryID  *int    `json:"category_id"`
	Description string  `json:"description" binding:"required"`
	ImageURL    *string `json:"image_url"`
	ReleaseYear *int    `json:"release_year"`
	Price       *int    `json:"price"`
	TotalPage   *int    `json:"total_page"`
	CreatedBy   *string `json:"created_by"`
	ModifiedBy  *string `json:"modified_by"`
}

type Buku struct {
	ID           int        `json:"id"`
	Title        string     `json:"title"`
	CategoryID   *int       `json:"category_id"`
	CategoryName *string    `json:"category_name,omitempty"`
	Description  string     `json:"description"`
	ImageURL     *string    `json:"image_url"`
	ReleaseYear  *int       `json:"release_year"`
	Price        *int       `json:"price"`
	TotalPage    *int       `json:"total_page"`
	Thickness    *string    `json:"thickness"`
	CreatedAt    time.Time  `json:"created_at"`
	CreatedBy    *string    `json:"created_by"`
	ModifiedAt   *time.Time `json:"modified_at"`
	ModifiedBy   *string    `json:"modified_by"`
}

type Users struct {
	ID         int        `json:"id"`
	Username   string     `json:"username" binding:"required"`
	Password   string     `json:"password" binding:"required"`
	CreatedAt  time.Time  `json:"created_at"`
	CreatedBy  *string    `json:"created_by"`
	ModifiedAt *time.Time `json:"modified_at"`
	ModifiedBy *string    `json:"modified_by"`
}

type Kategori struct {
	ID         int        `json:"id"`
	Name       string     `json:"name" binding:"required"`
	CreatedAt  time.Time  `json:"created_at"`
	CreatedBy  string     `json:"created_by"`
	ModifiedAt *time.Time `json:"modified_at"`
	ModifiedBy *string    `json:"modified_by"`
}
