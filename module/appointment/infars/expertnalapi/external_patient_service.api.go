package externalapi

type externalPatientService struct {
	apiURL string
}

func NewPatientServiceAPI(apiURL string) *externalPatientService {
	return &externalPatientService{apiURL: apiURL}
}
