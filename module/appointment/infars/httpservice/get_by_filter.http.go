package appointmenthttpservice

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	appointmentdomain "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/domain"
	appointmentqueries "github.com/PhuPhuoc/curanest-appointment-service/module/appointment/usecase/queries"
)

//	@Summary		get appointment by filter option
//	@Description	get appointment by filter option
//	@Tags			appointments
//	@Accept			json
//	@Produce		json
//	@Param			service-id			query		string					false	"service ID (UUID)"
//	@Param			cuspackage-id		query		string					false	"customized package ID (UUID)"
//	@Param			nursing-id			query		string					false	"nursing ID (UUID)"
//	@Param			patient-id			query		string					false	"patient ID (UUID)"
//	@Param			had-nurse			query		string					false	"had a nurse not not"
//	@Param			appointment-status	query		string					false	"appointment status"
//	@Param			est-date-from		query		string					false	"est date from (YYYY-MM-DD)"
//	@Param			est-date-to			query		string					false	"est date to (YYYY-MM-DD)"
//	@Success		200					{object}	map[string]interface{}	"data"
//	@Failure		400					{object}	error					"Bad request error"
//	@Router			/api/v1/appointments [get]
//	@Security		ApiKeyAuth
func (s *appointmentHttpService) handleGetAppointmentByFilter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		filter := &appointmentqueries.FilterGetAppointmentDTO{}

		if serviceId := ctx.Query("service-id"); serviceId != "" {
			serviceUUID, err := uuid.Parse(serviceId)
			if err != nil {
				common.ResponseError(ctx, common.NewBadRequestError().WithReason("service-id invalid (not a UUID)"))
				return
			}
			filter.ServiceId = &serviceUUID
		}

		if cusPackageId := ctx.Query("cuspackage-id"); cusPackageId != "" {
			cusPackageUUID, err := uuid.Parse(cusPackageId)
			if err != nil {
				common.ResponseError(ctx, common.NewBadRequestError().WithReason("cuspackage-id invalid (not a UUID)"))
				return
			}
			filter.CusPackageId = &cusPackageUUID
		}

		if nursingId := ctx.Query("nursing-id"); nursingId != "" {
			nursingUUID, err := uuid.Parse(nursingId)
			if err != nil {
				common.ResponseError(ctx, common.NewBadRequestError().WithReason("nursing-id invalid (not a UUID)"))
				return
			}
			filter.NursingId = &nursingUUID
		}

		if patientId := ctx.Query("patient-id"); patientId != "" {
			patientUUID, err := uuid.Parse(patientId)
			if err != nil {
				common.ResponseError(ctx, common.NewBadRequestError().WithReason("patient-id invalid (not a UUID)"))
				return
			}
			filter.PatientId = &patientUUID
		}

		if hadNurse := ctx.Query("had-nurse"); hadNurse != "" {
			hadNurseBool, err := strconv.ParseBool(hadNurse)
			if err != nil {
				common.ResponseError(ctx, common.NewBadRequestError().WithReason("had-nurse must be a bool"))
				return
			}
			filter.HadNurse = &hadNurseBool
		}

		if status := ctx.Query("appointment-status"); status != "" {
			appointmentStatus := appointmentdomain.EnumAppointmentStatus(status)
			filter.AppointmentStatus = &appointmentStatus
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

		appointments, err := s.query.GetAppointment.Handle(ctx, filter)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseGetWithPagination(ctx, appointments, nil, filter)
	}
}
