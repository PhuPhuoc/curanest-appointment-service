package categoryrepository

import "github.com/jmoiron/sqlx"

type categoryRepo struct {
	db *sqlx.DB
}

func NewCategoryRepo(db *sqlx.DB) *categoryRepo {
	return &categoryRepo{
		db: db,
	}
}
