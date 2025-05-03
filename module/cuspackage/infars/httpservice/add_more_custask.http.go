package cuspackagehttpservice

import (
	"github.com/PhuPhuoc/curanest-appointment-service/common"
	cuspackagecommands "github.com/PhuPhuoc/curanest-appointment-service/module/cuspackage/usecase/commands"
	"github.com/gin-gonic/gin"
)

// @Summary		add new customized service into a appointment (client)
// @Description	add new customized service into a appointment (client)
// @Tags			customized packages
// @Accept			json
// @Produce		json
// @Param			create	form		body					cuspackagecommands.AddMoreCustaskRequestDTO	true	"customized task creation data"
// @Success		200		{object}	map[string]interface{}	"data"
// @Failure		400		{object}	error					"Bad request error"
// @Router			/api/v1/cuspackage/add-more-custask [post]
// @Security		ApiKeyAuth
func (s *cusPackageHttpService) handleAddNewCustaskIntoApointment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto cuspackagecommands.AddMoreCustaskRequestDTO
		if err := ctx.BindJSON(&dto); err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("invalid request body").WithInner(err.Error()))
			return
		}

		err := s.cmd.AddNewCustaskIntoAppointment.Handle(ctx.Request.Context(), &dto)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseCreated(ctx)
	}
}
