package externalapi

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	appointmentqueries "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/usecase/queries"
	"github.com/google/uuid"
)

func (ex *externalNursingService) GetNursingByServiceIdRPC(ctx context.Context, serviceId uuid.UUID) ([]appointmentqueries.NurseDTO, error) {
	response, err := common.CallExternalAPI(ctx, common.RequestOptions{
		Method: "GET",
		URL:    ex.apiURL + "/external/rpc/nurses/service/" + serviceId.String(),
	})
	if err != nil {
		resp := common.NewInternalServerError().WithReason("cannot call external api: " + err.Error())
		return nil, resp
	}

	success, ok := response["success"].(bool)
	if !ok || !success {
		return nil, common.ExtractErrorFromResponse(response)
	}

	return common.ExtractListDataFromResponse[appointmentqueries.NurseDTO](response, "data")
}
