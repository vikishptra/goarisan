package arisanapi

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"vikishptra/domain_goarisan/usecase/runuserlogin"
	"vikishptra/shared/gogen"
	"vikishptra/shared/infrastructure/logger"
	"vikishptra/shared/model/payload"
	"vikishptra/shared/util"
)

func (r *ginController) runUserLoginHandler() gin.HandlerFunc {

	type InportRequest = runuserlogin.InportRequest
	type InportResponse = runuserlogin.InportResponse

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

		r.log.Info(ctx, util.MustJSON(req))

		res, err := inport.Execute(ctx, req)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		jsonRes.Name = res.Name
		jsonRes.Token = res.Token
		jsonRes.Password = res.Password
		jsonRes.Now = res.Now
		jsonRes.RandomString = res.RandomString

		r.log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
