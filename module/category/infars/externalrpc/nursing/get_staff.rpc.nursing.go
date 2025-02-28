package categorynursingrpc

import (
	"context"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	categoryqueries "github.com/PhuPhuoc/curanest-appointment-service/module/category/usecase/queries"
)

func (rpcService *externalNursingService) GetStaffsRPC(ctx context.Context, ids *categoryqueries.StaffIdsQueryDTO) ([]categoryqueries.StaffDTO, error) {
	token, ok := ctx.Value(common.KeyToken).(string)
	if !ok {
		return nil, common.NewInternalServerError().
			WithReason("cannot get accounts profile").WithInner("missing token to fetching data from other service")
	}

	response, err := common.CallExternalAPI(ctx, common.RequestOptions{
		Method:  "POST",
		URL:     rpcService.apiURL + "/external/rpc/nurses/by-ids",
		Token:   token,
		Payload: ids,
	})
	if err != nil {
		resp := common.NewInternalServerError().WithReason("cannot call external api: " + err.Error())
		return nil, resp
	}

	success, ok := response["success"].(bool)
	if !ok || !success {
		return nil, common.ExtractErrorFromResponse(response)
	}

	return common.ExtractListDataFromResponse[categoryqueries.StaffDTO](response, "data")
}
