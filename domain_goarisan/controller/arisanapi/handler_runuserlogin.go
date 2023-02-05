package arisanapi

import (
	"context"
	"net/http"
	"os"

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
		req.Email = jsonReq.Email
		req.Password = jsonReq.Password

		r.log.Info(ctx, util.MustJSON(req))

		res, err := inport.Execute(ctx, req)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		jsonRes.Email = res.Email
		jsonRes.Token = res.Token
		jsonRes.Now = res.Now
		jsonRes.RandomString = res.RandomString
		jsonRes.Name = res.Name
		jsonRes.RefreshToken = res.RefreshToken
		domain := os.Getenv("ORIGIN")
		timeToken := 24 * 60 * 60 * 100
		c.SetCookie("refresh_token", res.RefreshToken, timeToken, "/", domain, false, true)
		r.log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
