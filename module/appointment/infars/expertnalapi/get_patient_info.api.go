package externalapi

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	apppointmentcommands "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/usecase/commands"
	"github.com/google/uuid"
)

func (ex *externalPatientService) GetPatientInfo(ctx context.Context, patientId uuid.UUID) (*apppointmentcommands.PatientInfo, error) {
	response, err := common.CallExternalAPI(ctx, common.RequestOptions{
		Method: "GET",
		URL:    ex.apiURL + "/api/v1/patients/" + patientId.String(),
	})
	if err != nil {
		resp := common.NewInternalServerError().WithReason("cannot call external api: " + err.Error())
		return nil, resp
	}

	success, ok := response["success"].(bool)
	if !ok || !success {
		return nil, common.ExtractErrorFromResponse(response)
	}

	return common.ExtractDataFromResponse[apppointmentcommands.PatientInfo](response, "data")
}
