package externalapi

type externalPushNotiService struct {
	apiURL string
}

func NewPushNotiServiceRPC(apiURL string) *externalPushNotiService {
	return &externalPushNotiService{apiURL: apiURL}
}
