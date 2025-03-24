package cuspackagecommands

import "context"

type createCusPackageAndTaskHandler struct {
	cmdRepo CusPackageCommandRepo
}

func NewCreateCusPackageAndTaskHandler(cmdRepo CusPackageCommandRepo) *createCusPackageAndTaskHandler {
	return &createCusPackageAndTaskHandler{
		cmdRepo: cmdRepo,
	}
}

func (h *createCusPackageAndTaskHandler) Handle(ctx context.Context, cuspack *CreateCustomizedPackageDTO, custask []CreateCustomizedTaskDTO) error {
	return nil
}
