package appointmenthttpservice

import (
	"time"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	appointmentqueries "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/usecase/queries"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//	@Summary		get timesheet of nursing
//	@Description	get timesheet of nursing
//	@Tags			appointments
//	@Accept			json
//	@Produce		json
//	@Param			nursing-id		query		string					false	"nursing ID (UUID)"
//	@Param			est-date-from	query		string					false	"est date from (YYYY-MM-DD)"
//	@Param			est-date-to		query		string					false	"est date to (YYYY-MM-DD)"
//	@Success		200				{object}	map[string]interface{}	"data"
//	@Failure		400				{object}	error					"Bad request error"
//	@Router			/api/v1/appointments/nursing-timesheet [get]
//	@Security		ApiKeyAuth
func (s *appointmentHttpService) handleGetNursingTimesheet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		filter := &appointmentqueries.FilterGetNursingTimesheetDTO{}

		if nursingId := ctx.Query("nursing-id"); nursingId != "" {
			nursingUUID, err := uuid.Parse(nursingId)
			if err != nil {
				common.ResponseError(ctx, common.NewBadRequestError().WithReason("nursing-id invalid (not a UUID)"))
				return
			}
			filter.NursingId = nursingUUID
		}

		if estDateFrom := ctx.Query("est-date-from"); estDateFrom != "" {
			parsedDate, err := time.Parse("2006-01-02", estDateFrom)
			if err != nil {
				common.ResponseError(ctx, common.NewBadRequestError().WithReason("est-date-from invalid (use YYYY-MM-DD)"))
				return
			}
			filter.EstDateFrom = &parsedDate
		}

		if estDateTo := ctx.Query("est-date-to"); estDateTo != "" {
			parsedDate, err := time.Parse("2006-01-02", estDateTo)
			if err != nil {
				common.ResponseError(ctx, common.NewBadRequestError().WithReason("est-date-to invalid (use YYYY-MM-DD)"))
				return
			}
			filter.EstDateTo = &parsedDate
		}

		timesheets, err := s.query.GetNursingTimeSheet.Handle(ctx, filter)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseSuccess(ctx, timesheets)
	}
}
