package externalapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	apppointmentcommands "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/usecase/commands"
)

func (ex *externalGoongAPI) GetDistanceFromGoong(ctx context.Context, originCode, destinationCode string) (*apppointmentcommands.DistanceMatrixResponse, error) {
	apiURL := fmt.Sprintf("https://rsapi.goong.io/DistanceMatrix?api_key=%s&origins=%s&destinations=%s&vehicle=hd", ex.apiKey, originCode, destinationCode)

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to call API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API returned non-200 status: %d, body: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var apiResponse *apppointmentcommands.DistanceMatrixResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v, body: %s", err, string(body))
	}

	if apiResponse.Rows[0].Elements[0].Status != "OK" {
		return nil, fmt.Errorf("API returned non-OK status: %v", apiResponse.Rows[0].Elements[0].Status != "OK")
	}

	return apiResponse, nil
}
