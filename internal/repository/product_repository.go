package repository

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/alviankristi/catalyst-backend-task/internal/repository/entity"
	"github.com/alviankristi/catalyst-backend-task/pkg/response"
)

var (
	createProduct    string = `INSERT INTO products (name, brand_id, price, created_date) VALUES (?, ?, ?, ?)`
	getSingleProduct string = `SELECT id, created_date, brand_id ,name, price FROM products where id = ?`
	getByBrandId     string = `SELECT id, created_date, brand_id ,name, price FROM products where brand_id = ?`
	productNameExist string = `SELECT 1 FROM products where LOWER(name) = LOWER(?) and brand_id = ?`
	productPrice     string = `SELECT price FROM products where id = ?`
)

type ProductRepository interface {
	Create(context context.Context, model entity.CreatedProductEntity) (*entity.ProductEntity, error)
	GetById(context context.Context, id int) (*entity.ProductEntity, error)
	GetByBrandId(context context.Context, id int) ([]*entity.ProductEntity, error)
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (repo *productRepository) GetById(context context.Context, id int) (*entity.ProductEntity, error) {
	result := &entity.ProductEntity{}
	err := repo.db.QueryRowContext(context, getSingleProduct, id).
		Scan(&result.Id, &result.CreatedDate, &result.BrandId, &result.Name, &result.Price)
	switch err {
	case nil:
		//return entity
		return result, nil
	case sql.ErrNoRows:
		//name not exist
		return nil, nil
	default:
		//error db
		log.Printf("productRepository.GetById() error : %v", err)
		return nil, response.DatabaseError
	}
}

func (repo *productRepository) GetByBrandId(context context.Context, id int) ([]*entity.ProductEntity, error) {
	rows, err := repo.db.QueryContext(context, getByBrandId, id)
	if err != nil {
		log.Printf("productRepository.GetByBrandId() error : %v", err)
		return nil, response.DatabaseError
	}

	defer rows.Close()

	results := []*entity.ProductEntity{}
	for rows.Next() {
		result := &entity.ProductEntity{}
		err := rows.Scan(&result.Id, &result.CreatedDate, &result.BrandId, &result.Name, &result.Price)
		if err != nil {
			log.Printf("productRepository.GetByBrandId() - rows.Scan() error : %v", err)
		} else {
			results = append(results, result)
		}
	}

	err = rows.Err()
	if err != nil {
		log.Printf("productRepository.GetByBrandId() - rows.Err() error : %v", err)
		return nil, response.DatabaseError
	}
	return results, nil
}

func (repo *productRepository) Create(context context.Context, model entity.CreatedProductEntity) (*entity.ProductEntity, error) {
	//check name duplicate or not
	if err := repo.validateBrandProductNameDuplicate(context, model); err != nil {
		return nil, err
	}

	//save db
	createdDate := time.Now()
	result, err := repo.db.ExecContext(
		context,
		createProduct,
		model.Name, model.BrandId, model.Price, createdDate)

	if err != nil {
		log.Printf("productRepository.Create() - repo.db.ExecContext() error : %v", err)
		return nil, response.DatabaseError
	}

	//get id
	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("productRepository.Create() - result.LastInsertId() error : %v", err)
		return nil, response.DatabaseError
	}

	m := &entity.ProductEntity{
		Name:       model.Name,
		BrandId:    model.BrandId,
		Price:      model.Price,
		BaseEntity: entity.NewBaseEntity(id, createdDate),
	}
	return m, nil
}

func (repo *productRepository) validateBrandProductNameDuplicate(ctx context.Context, model entity.CreatedProductEntity) error {

	var result *int
	err := repo.db.QueryRowContext(ctx, productNameExist, model.Name, model.BrandId).Scan(&result)

	switch err {
	case nil:
		//name already exist
		log.Printf("productRepository.validateBrandProductNameDuplicate() error : %v", response.BrandProductNameDuplicate)
		return response.BrandProductNameDuplicate
	case sql.ErrNoRows:
		//name not exist
		return nil
	default:
		//error db
		log.Printf("productRepository.validateBrandProductNameDuplicate() error : %v", err)
		return response.DatabaseError
	}
}
