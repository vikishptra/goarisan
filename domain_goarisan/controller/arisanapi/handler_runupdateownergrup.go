package arisanapi

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"vikishptra/domain_goarisan/model/errorenum"
	"vikishptra/domain_goarisan/usecase/runupdateownergrup"
	"vikishptra/shared/gogen"
	"vikishptra/shared/infrastructure/logger"
	"vikishptra/shared/infrastructure/token"
	"vikishptra/shared/model/payload"
	"vikishptra/shared/util"
)

func (r *ginController) runUpdateOwnerGrupHandler() gin.HandlerFunc {

	type InportRequest = runupdateownergrup.InportRequest
	type InportResponse = runupdateownergrup.InportResponse

	inport := gogen.GetInport[InportRequest, InportResponse](r.GetUsecase(InportRequest{}))

	type request struct {
		InportRequest
	}

	type response struct {
		InportResponse
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID()

		ctx := logger.SetTraceID(context.Background(), traceID)
		id, _ := token.ExtractTokenID(c)

		var jsonReq request
		if err := c.BindJSON(&jsonReq); err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var req InportRequest
		req.IDUser = jsonReq.IDUser
		req.IDGrup = jsonReq.IDGrup
		req.IDOwner = id

		r.log.Info(ctx, util.MustJSON(req))

		res, err := inport.Execute(ctx, req)
		if err != nil {
			if err == errorenum.AndaBukanAdmin {
				r.log.Error(ctx, err.Error())
				c.JSON(http.StatusForbidden, payload.NewErrorResponse(err, traceID))
				return
			} else if err == errorenum.DataUserNotFound {
				r.log.Error(ctx, err.Error())
				c.JSON(http.StatusNotFound, payload.NewErrorResponse(err, traceID))
				return
			}
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		jsonRes.Message = res.Message

		r.log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
