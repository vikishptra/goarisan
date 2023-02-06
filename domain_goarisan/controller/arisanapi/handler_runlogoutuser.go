package arisanapi

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"vikishptra/domain_goarisan/model/errorenum"
	"vikishptra/domain_goarisan/usecase/runlogoutuser"
	"vikishptra/shared/gogen"
	"vikishptra/shared/infrastructure/logger"
	"vikishptra/shared/infrastructure/token"
	"vikishptra/shared/model/payload"
	"vikishptra/shared/util"
)

func (r *ginController) runLogoutUserHandler() gin.HandlerFunc {

	type InportRequest = runlogoutuser.InportRequest
	type InportResponse = runlogoutuser.InportResponse

	inport := gogen.GetInport[InportRequest, InportResponse](r.GetUsecase(InportRequest{}))

	type response struct {
		InportResponse
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID()

		ctx := logger.SetTraceID(context.Background(), traceID)

		var req InportRequest
		id, _ := token.ExtractTokenID(c)
		refreshToken := token.ExtractTokenCookie(c)
		if refreshToken == "" {
			c.JSON(http.StatusUnauthorized, payload.NewErrorResponse(errorenum.GabisaAksesBro, traceID))
			return
		}
		req.Token = id

		r.log.Info(ctx, util.MustJSON(req))
		res, err := inport.Execute(ctx, req)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}
		//ubah token menjadi nulll
		domain := os.Getenv("DOMAIN")
		c.SetCookie("refresh_token", "", time.Now().Second(), "/", domain, false, true)

		var jsonRes response
		jsonRes.Message = res.Message
		r.log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
