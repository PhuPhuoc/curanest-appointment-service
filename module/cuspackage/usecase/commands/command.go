package cuspackagecommands

type Commands struct{}

type Builder interface {
	BuildCusPackageCmdRepo() CusPackageCommandRepo
}

func NewCusPackageCmdWithBuilder(b Builder) Commands {
	return Commands{}
}

type CusPackageCommandRepo interface{}
