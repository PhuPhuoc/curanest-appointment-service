package cuspackagerepository

import "github.com/jmoiron/sqlx"

type cusPackageRepo struct {
	db *sqlx.DB
}

func NewCusPackageRepo(db *sqlx.DB) *cusPackageRepo {
	return &cusPackageRepo{
		db: db,
	}
}
