package entity

type ProductEntity struct {
	BaseEntity
	Name    string `db:"name"`
	BrandId int    `db:"brand_id"`
	Price   int    `db:"price"`
}

type CreatedProductEntity struct {
	Name    string `db:"name"`
	BrandId int    `db:"brand_id"`
	Price   int    `db:"price"`
}
