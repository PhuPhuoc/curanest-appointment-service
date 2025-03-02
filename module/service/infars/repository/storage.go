package servicerepository

import "github.com/jmoiron/sqlx"

type serviceRepo struct {
	db *sqlx.DB
}

func NewServiceRepo(db *sqlx.DB) *serviceRepo {
	return &serviceRepo{
		db: db,
	}
}
