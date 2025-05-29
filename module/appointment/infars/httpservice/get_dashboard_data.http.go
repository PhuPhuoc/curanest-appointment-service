package appointmenthttpservice

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	appointmentqueries "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/usecase/queries"
)

// @Summary		get appointment by filter option
// @Description	get appointment by filter option
// @Tags			appointments
// @Accept			json
// @Produce		json
// @Param			is-admin	query		string					false	"is admin or staff"
// @Param			category-id	query		string					false	"category ID (UUID)"
// @Param			date-from	query		string					false	"date from (YYYY-MM-DD)"
// @Param			date-to		query		string					false	"date to (YYYY-MM-DD)"
// @Success		200			{object}	map[string]interface{}	"data"
// @Failure		400			{object}	error					"Bad request error"
// @Router			/api/v1/appointments/dashboard  [get]
// @Security		ApiKeyAuth
func (s *appointmentHttpService) handleGetDashboardData() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		filter := &appointmentqueries.FilterDashboardDTO{}

		if isAdmin := ctx.Query("is-admin"); isAdmin != "" {
			isAdminBool, err := strconv.ParseBool(isAdmin)
			if err != nil {
				common.ResponseError(ctx, common.NewBadRequestError().WithReason("had-nurse must be a bool"))
				return
			}
			filter.IsAdmin = isAdminBool
		}

		if categoryId := ctx.Query("category-id"); categoryId != "" {
			cateUUID, err := uuid.Parse(categoryId)
			if err != nil {
				common.ResponseError(ctx, common.NewBadRequestError().WithReason("category-id invalid (not a UUID)"))
				return
			}
			filter.CategoryId = &cateUUID
		}

		if dateFrom := ctx.Query("date-from"); dateFrom != "" {
			// parsedDate, err := time.Parse("2006-01-02", dateFrom)
			// if err != nil {
			// 	common.ResponseError(ctx, common.NewBadRequestError().WithReason("date-from invalid (use YYYY-MM-DD)"))
			// 	return
			// }
			filter.DateFrom = dateFrom
		}

		if dateTo := ctx.Query("date-to"); dateTo != "" {
			// parsedDate, err := time.Parse("2006-01-02", dateTo)
			// if err != nil {
			// 	common.ResponseError(ctx, common.NewBadRequestError().WithReason("est-date-to invalid (use YYYY-MM-DD)"))
			// 	return
			// }
			filter.DateTo = dateTo
		}

		resp, err := s.query.GetDashboardData.Handle(ctx, filter)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}
		common.ResponseSuccess(ctx, resp)
	}
}
