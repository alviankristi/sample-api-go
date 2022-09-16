package repository

import "database/sql"

type Repository struct {
	BrandRepository BrandRepository
}

//create repository
func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		BrandRepository: NewBrandRepository(db),
	}
}
