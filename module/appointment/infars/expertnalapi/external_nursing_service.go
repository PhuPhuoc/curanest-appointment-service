package externalapi

type externalNursingService struct {
	apiURL string
}

func NewNursingRPC(apiURL string) *externalNursingService {
	return &externalNursingService{apiURL: apiURL}
}
