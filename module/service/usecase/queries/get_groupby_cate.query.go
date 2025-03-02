package servicequeries

import "context"

type getServicesGroupByCategoryHandler struct {
	queryRepo   ServiceQueryRepo
	cateFetcher CategoryFetcher
}

func NewGetServicesGroupByCategoryHandler(queryRepo ServiceQueryRepo, cateFetcher CategoryFetcher) *getServicesGroupByCategoryHandler {
	return &getServicesGroupByCategoryHandler{
		queryRepo:   queryRepo,
		cateFetcher: cateFetcher,
	}
}

func (h *getServicesGroupByCategoryHandler) Handle(ctx context.Context, filter FilterGetService) ([]ServiceDTO, error) {
	return nil, nil
}
