package arisanapi

import (
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"vikishptra/domain_goarisan/model/errorenum"
	"vikishptra/domain_goarisan/usecase/runverifynewpasswordemail"
	"vikishptra/shared/gogen"
	"vikishptra/shared/infrastructure/logger"
	"vikishptra/shared/model/payload"
	"vikishptra/shared/util"
)

func (r *ginController) runverifynewpasswordemailHandler() gin.HandlerFunc {

	type InportRequest = runverifynewpasswordemail.InportRequest
	type InportResponse = runverifynewpasswordemail.InportResponse

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
		if err := c.Bind(&jsonReq); err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var req InportRequest
		req.Code = jsonReq.Code
		req.Id = jsonReq.Id

		r.log.Info(ctx, util.MustJSON(req))

		res, err := inport.Execute(ctx, req)
		if err != nil {
			if err == errorenum.TokenAndaSudahKadaluawarsa {
				c.Redirect(http.StatusFound, os.Getenv("DOMAIN_REDIRECT")+"/tokenkadaluarsa")
			}
			r.log.Error(ctx, err.Error())
			c.Redirect(http.StatusFound, os.Getenv("DOMAIN_REDIRECT")+"/kesalahanuser")
			return
		}

		var jsonRes response
		jsonRes.Code = res.Code
		jsonRes.Id = res.Id
		r.log.Info(ctx, util.MustJSON(jsonRes))
		url := os.Getenv("DOMAIN_REDIRECT") + "/new/password?token=" + jsonRes.Code + "&id=" + jsonRes.Id
		c.Redirect(http.StatusFound, url)
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
