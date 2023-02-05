package arisanapi

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"vikishptra/domain_goarisan/usecase/runusercreate"
	"vikishptra/shared/gogen"
	"vikishptra/shared/infrastructure/logger"
	"vikishptra/shared/model/payload"
	"vikishptra/shared/util"
)

func (r *ginController) runUserCreateHandler() gin.HandlerFunc {

	type InportRequest = runusercreate.InportRequest
	type InportResponse = runusercreate.InportResponse

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

		var jsonReq request
		if err := c.BindJSON(&jsonReq); err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var req InportRequest
		req.Name = jsonReq.Name
		req.Password = jsonReq.Password
		req.Email = jsonReq.Email
		req.ConfirmPassword = jsonReq.ConfirmPassword
		req.Now = time.Now()
		req.RandomString = util.GenerateID()

		r.log.Info(ctx, util.MustJSON(req))

		res, err := inport.Execute(ctx, req)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		jsonRes.Message = res.Message
		r.log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusCreated, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
