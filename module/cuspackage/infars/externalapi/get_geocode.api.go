package externalapigoong

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	cuspackagecommands "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/usecase/commands"
)

func (ex *externalGoongAPI) GetGeocodeFromGoong(ctx context.Context, address string) (*cuspackagecommands.GoongAPIResponse, error) {
	response, err := common.CallExternalAPI(ctx, common.RequestOptions{
		Method: "GET",
		URL:    ex.apiURL + "/geocode?address=" + address + "&api_key=" + ex.apiKey,
	})
	if err != nil {
		return nil, common.NewInternalServerError().
			WithReason("cannot get api get geocode from goong").WithInner("cannot call external api - " + err.Error())
	}

	success, ok := response["status"].(string)
	if !ok || success != "OK" {
		return nil, common.ExtractErrorFromResponse(response)
	}

	return common.ExtractDataFromResponse[cuspackagecommands.GoongAPIResponse](response, "results")
}
