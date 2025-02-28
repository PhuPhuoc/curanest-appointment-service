package categoryauthrpc

type externalAccountService struct {
	apiURL string
}

func NewAccountRPC(apiURL string) *externalAccountService {
	return &externalAccountService{apiURL: apiURL}
}
