package entity

import "time"

type BaseEntity struct {
	Id          int       `db:"id"`
	CreatedDate time.Time `db:"created_date"`
}

func NewBaseEntity(id int64, createdDate time.Time) BaseEntity {
	return BaseEntity{
		Id:          int(id),
		CreatedDate: createdDate,
	}
}
