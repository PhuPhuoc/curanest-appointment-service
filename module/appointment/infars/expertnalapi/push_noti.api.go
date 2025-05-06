package externalapi

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
)

func (ex *externalPushNotiService) PushNotification(ctx context.Context, req *common.PushNotiRequest) error {
	// Gọi external API
	result, err := common.CallExternalAPI(ctx, common.RequestOptions{
		Method:  "POST",
		URL:     ex.apiURL + "/external/rpc/notifications",
		Payload: req,
	})
	if err != nil {
		return common.NewInternalServerError().WithReason("cannot call external api: " + err.Error())
	}

	// Chuyển map kết quả về JSON string
	rawJSON, err := json.Marshal(result)
	if err != nil {
		return fmt.Errorf("failed to marshal result: %v", err)
	}

	var parsed common.PushNotiResponse
	if err := json.Unmarshal(rawJSON, &parsed); err != nil {
		return fmt.Errorf("failed to unmarshal response: %v", err)
	}

	if !parsed.Success {
		return common.NewBadRequestError().WithReason(parsed.Message)
	}

	return nil
}
