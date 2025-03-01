package categoryauthrpc

import (
	"context"

	"github.com/google/uuid"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
)

func (rpcSer *externalAccountService) UpdateAccountRoleRPC(ctx context.Context, nursingId uuid.UUID, targetRole string) error {
	if targetRole == "" || targetRole != "staff" && targetRole != "nurse" {
		return common.NewBadRequestError().
			WithReason("transfer between staff and nurse failed").WithInner("request role must be 'staff' or 'nurse'")
	}

	token, ok := ctx.Value(common.KeyToken).(string)
	if !ok {
		return common.NewInternalServerError().
			WithReason("cannot get accounts profile").WithInner("missing token to fetching data from other service")
	}

	response, err := common.CallExternalAPI(ctx, common.RequestOptions{
		Method: "PATCH",
		URL:    rpcSer.apiURL + "/external/rpc/accounts/" + nursingId.String() + "/role?target-role=" + targetRole,
		Token:  token,
	})
	if err != nil {
		resp := common.NewInternalServerError().WithReason("cannot call external api: " + err.Error())
		return resp
	}

	success, ok := response["success"].(bool)
	if !ok || !success {
		return common.ExtractErrorFromResponse(response)
	}

	return nil
}
