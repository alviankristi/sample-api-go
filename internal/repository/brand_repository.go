package repository

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/alviankristi/catalyst-backend-task/internal/repository/entity"
	"github.com/alviankristi/catalyst-backend-task/pkg/response"
)

type BrandRepository interface {
	Create(ctx context.Context, name string) (*entity.BrandEntity, error)
}

type brandRepository struct {
	db *sql.DB
}

var (
	createBrand    string = `INSERT INTO brands (name, created_date) VALUES (?,?)`
	brandNameExist string = `SELECT 1 FROM brands where LOWER(name) = LOWER(?)`
)

func NewBrandRepository(db *sql.DB) BrandRepository {
	return &brandRepository{
		db: db,
	}
}

func (repo brandRepository) Create(ctx context.Context, name string) (*entity.BrandEntity, error) {

	if err := repo.validateBrandNameDuplicate(ctx, name); err != nil {
		return nil, err
	}

	createdDate := time.Now()
	result, err := repo.db.ExecContext(
		ctx,
		createBrand,
		name, createdDate)

	if err != nil {
		log.Printf("brandRepository.Create() - repo.db.ExecContext() error : %v", err)
		return nil, response.DatabaseError
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("brandRepository.Create() - result.LastInsertId() error : %v", err)
		return nil, response.DatabaseError
	}

	m := &entity.BrandEntity{
		Name:       name,
		BaseEntity: entity.NewBaseEntity(id, createdDate),
	}
	return m, nil
}

func (repo brandRepository) validateBrandNameDuplicate(ctx context.Context, name string) error {

	var result *int
	err := repo.db.QueryRowContext(ctx, brandNameExist, name).Scan(&result)

	switch err {
	case nil:
		log.Printf("brandRepository.validateBrandNameDuplicate() error : %v", response.BrandNameDuplicate)
		return response.BrandNameDuplicate
	case sql.ErrNoRows:
		return nil
	default:
		log.Printf("brandRepository.validateBrandNameDuplicate() error : %v", err)
		return response.DatabaseError
	}
}
