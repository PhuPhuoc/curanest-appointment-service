package servicecommands

type CreateServiceDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	EstDuration string `json:"est-duration"`
}
