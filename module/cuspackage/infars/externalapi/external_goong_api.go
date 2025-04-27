package externalapi

type externalGoongAPI struct {
	apiURL string
	apiKey string
}

func NewExternalGoongAPI(apiURL, apiKey string) *externalGoongAPI {
	return &externalGoongAPI{
		apiURL: apiURL,
		apiKey: apiKey,
	}
}
