package svcpackagerepository

import "github.com/jmoiron/sqlx"

type svcPackageRepo struct {
	db *sqlx.DB
}

func NewSvcPackageRepo(db *sqlx.DB) *svcPackageRepo {
	return &svcPackageRepo{
		db: db,
	}
}
