package servicehttpservice

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/PhuPhuoc/curanest-appointment-service/common"
	servicecommands "github.com/PhuPhuoc/curanest-appointment-service/module/service/usecase/commands"
)

//	@Summary		create new service
//	@Description	create new service
//	@Tags			services
//	@Accept			json
//	@Produce		json
//	@Param			category-id	path		string					true								"category ID (UUID)"
//	@Param			create		form		body					servicecommands.CreateServiceDTO	true	"service creation data"
//	@Success		200			{object}	map[string]interface{}	"data"
//	@Failure		400			{object}	error					"Bad request error"
//	@Router			/api/v1/categories/{category-id}/services [post]
//	@Security		ApiKeyAuth
func (s *serviceHttpService) handleCreateService() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cateId := ctx.Param("category-id")
		if cateId == "" {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("missing category-id"))
			return
		}
		cateUUID, err := uuid.Parse(cateId)
		if err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("category-id invalid (not a uuid)"))
			return
		}

		var dto servicecommands.CreateServiceDTO
		if err := ctx.BindJSON(&dto); err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("invalid request body").WithInner(err.Error()))
			return
		}

		if err := s.command.CreateService.Handle(ctx.Request.Context(), cateUUID, &dto); err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseCreated(ctx)
	}
}
