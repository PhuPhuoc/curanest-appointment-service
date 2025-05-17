package externalapi

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	"github.com/google/uuid"
)

func (ex *externalPatientService) GetRelativesId(ctx context.Context, patientId uuid.UUID) (*uuid.UUID, error) {
	response, err := common.CallExternalAPI(ctx, common.RequestOptions{
		Method: "GET",
		URL:    ex.apiURL + "/api/v1/patients/" + patientId.String() + "/relatives-id",
	})
	if err != nil {
		resp := common.NewInternalServerError().WithReason("cannot call external api: " + err.Error())
		return nil, resp
	}

	success, ok := response["success"].(bool)
	if !ok || !success {
		return nil, common.ExtractErrorFromResponse(response)
	}

	data, ok := response["data"].(map[string]interface{})
	if !ok {
		return nil, common.NewInternalServerError().WithReason("invalid data format")
	}

	relativesIdStr, ok := data["relatives-id"].(string)
	if !ok {
		return nil, common.NewInternalServerError().WithReason("relatives-id not found or not a string")
	}

	relativesId, err := uuid.Parse(relativesIdStr)
	if err != nil {
		return nil, common.NewInternalServerError().WithReason("invalid uuid format: " + err.Error())
	}

	return &relativesId, nil
}
