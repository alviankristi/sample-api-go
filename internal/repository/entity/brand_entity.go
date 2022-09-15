package entity

type BrandEntity struct {
	BaseEntity
	Name string `db:"name"`
}
