package common

type PayOSConfig struct {
	ClientId    string
	ApiKey      string
	CheckSumKey string
}

func NewPayOs(clientId string, apiKey string, checkSumKey string) *PayOSConfig {
	return &PayOSConfig{
		ClientId:    clientId,
		ApiKey:      apiKey,
		CheckSumKey: checkSumKey,
	}
}
