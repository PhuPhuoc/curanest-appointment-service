package servicehttpservice

import (
	"github.com/gin-gonic/gin"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	servicecommands "github.com/PhuPhuoc/curanest-appointment-service/module/service/usecase/commands"
)

//	@Summary		create new service
//	@Description	create new service
//	@Tags			services
//	@Accept			json
//	@Produce		json
//	@Param			create	form		body					servicecommands.CreateServiceDTO	true	"service creation data"
//	@Success		200		{object}	map[string]interface{}	"data"
//	@Failure		400		{object}	error					"Bad request error"
//	@Router			/api/v1/services [post]
//	@Security		ApiKeyAuth
func (s *serviceHttpService) handleCreateService() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto servicecommands.CreateServiceDTO
		if err := ctx.BindJSON(&dto); err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("invalid request body").WithInner(err.Error()))
			return
		}

		if err := s.command.CreateService.Handle(ctx.Request.Context(), &dto); err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseCreated(ctx)
	}
}
