package svcpackagequeries

type Queries struct{}

type Builder interface {
	BuildSvcPackageQueryRepo() SvcPackageQueryRepo
}

func NewSvcPackageQueryWithBuilder(b Builder) Queries {
	return Queries{}
}

type SvcPackageQueryRepo interface{}
