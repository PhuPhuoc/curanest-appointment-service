package externalapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	cuspackagecommands "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/usecase/commands"
)

func (ex *externalGoongAPI) GetGeocodeFromGoong(ctx context.Context, address string) (*cuspackagecommands.GoongAPIResponse, error) {
	encodedAddress := url.QueryEscape(address)
	apiURL := fmt.Sprintf("https://rsapi.goong.io/geocode?address=%s&api_key=%s", encodedAddress, ex.apiKey)

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

	var apiResponse *cuspackagecommands.GoongAPIResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v, body: %s", err, string(body))
	}

	if apiResponse.Status != "OK" {
		return nil, fmt.Errorf("API returned non-OK status: %s", apiResponse.Status)
	}

	return apiResponse, nil
}
