package cuspackagequeries

type Queries struct{}

type Builder interface {
	BuildCusPackageQueryRepo() CusPackageQueryRepo
}

func NewCusPackageQueryWithBuilder(b Builder) Queries {
	return Queries{}
}

type CusPackageQueryRepo interface{}
