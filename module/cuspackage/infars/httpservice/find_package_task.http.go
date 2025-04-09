package cuspackagehttpservice

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	cuspackagequeries "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/usecase/queries"
)

// @Summary		find customized-package & customized-tasks with Id and est-date
// @Description	find customized-package & customized-tasks with Id and est-date
// @Tags			customized packages
// @Accept			json
// @Produce		json
// @Param			cus-package-id	query		string					true	"customized-package ID (UUID)"
// @Param			est-date		query		string					true	"est date (YYYY-MM-DD)"
// @Success		200				{object}	map[string]interface{}	"data"
// @Failure		400				{object}	error					"Bad request error"
// @Router			/api/v1/cuspackage [get]
// @Security		ApiKeyAuth
func (s *cusPackageHttpService) handleFindCusPackageTask() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		filter := &cuspackagequeries.FilterGetCusPackageTaskDTO{}

		if packageId := ctx.Query("cus-package-id"); packageId != "" {
			packageUUID, err := uuid.Parse(packageId)
			if err != nil {
				common.ResponseError(ctx, common.NewBadRequestError().WithReason("cus-package-id invalid (not a UUID)"))
				return
			}
			filter.CusPackageId = packageUUID
		} else {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("cus-package-id is required"))
			return
		}

		if estDate := ctx.Query("est-date"); estDate != "" {
			parsedDate, err := time.Parse(time.RFC3339, estDate)
			if err != nil {
				common.ResponseError(ctx, common.NewBadRequestError().WithReason("est-date invalid (use YYYY-MM-DD)"))
				return
			}
			filter.EstDate = parsedDate
		} else {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("est-date is required"))
			return
		}

		response, err := s.query.FindCusPackageTask.Handle(ctx, filter)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseSuccess(ctx, response)
	}
}
