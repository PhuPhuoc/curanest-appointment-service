package servicehttpservice

import (
	"github.com/PhuPhuoc/curanest-appointment-service/common"
	servicecommands "github.com/PhuPhuoc/curanest-appointment-service/module/service/usecase/commands"
	"github.com/gin-gonic/gin"
)

//	@Summary		create new service
//	@Description	create new service
//	@Tags			services
//	@Accept			json
//	@Produce		json
//	@Param			update	form		body					servicecommands.UpdateServiceDTO	true	"service creation data"
//	@Success		200		{object}	map[string]interface{}	"data"
//	@Failure		400		{object}	error					"Bad request error"
//	@Router			/api/v1/services [put]
//	@Security		ApiKeyAuth
func (s *serviceHttpService) handleUpdateService() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto servicecommands.UpdateServiceDTO
		if err := ctx.BindJSON(&dto); err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("invalid request body").WithInner(err.Error()))
			return
		}

		if err := s.command.UpdateService.Handle(ctx.Request.Context(), dto); err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseCreated(ctx)
	}
}
