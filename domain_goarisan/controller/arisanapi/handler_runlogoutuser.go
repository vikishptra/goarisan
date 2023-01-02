package arisanapi

import (
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"vikishptra/domain_goarisan/controller/arisanapi/token"
	"vikishptra/domain_goarisan/usecase/runlogoutuser"
	"vikishptra/shared/gogen"
	"vikishptra/shared/infrastructure/logger"
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

		r.log.Info(ctx, util.MustJSON(req))
		res, err := inport.Execute(ctx, req)
		if err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var access_token string
		getAuth := token.ExtractToken(c)
		cookie, _ := c.Cookie("token")
		access_token = cookie
		if access_token != getAuth || access_token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Anda belum login"})
			return
		}
		domain := os.Getenv("DOMAIN")
		c.SetCookie("token", "", -1, "/", domain, false, true)
		c.SetCookie("logged_in", "", -1, "/", domain, false, true)

		var jsonRes response
		jsonRes.Message = res.Message
		r.log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
