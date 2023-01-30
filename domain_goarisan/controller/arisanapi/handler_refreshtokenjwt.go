package arisanapi

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"vikishptra/domain_goarisan/model/errorenum"
	"vikishptra/domain_goarisan/usecase/refreshtokenjwt"
	"vikishptra/shared/gogen"
	"vikishptra/shared/infrastructure/logger"
	"vikishptra/shared/infrastructure/token"
	"vikishptra/shared/model/payload"
	"vikishptra/shared/util"
)

func (r *ginController) refreshtokenjwtHandler() gin.HandlerFunc {

	type InportRequest = refreshtokenjwt.InportRequest
	type InportResponse = refreshtokenjwt.InportResponse

	inport := gogen.GetInport[InportRequest, InportResponse](r.GetUsecase(InportRequest{}))

	type response struct {
		InportResponse
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID()

		ctx := logger.SetTraceID(context.Background(), traceID)
		id, _ := token.ExtractTokenIDCookies(c)
		refreshToken := token.ExtractTokenCookie(c)
		var req InportRequest
		req.IDUser = id
		req.Token = refreshToken
		r.log.Info(ctx, util.MustJSON(req))

		res, err := inport.Execute(ctx, req)
		if err != nil {
			if err == errorenum.HayoMauNgapain {
				r.log.Error(ctx, err.Error())
				c.JSON(http.StatusForbidden, payload.NewErrorResponse(err, traceID))
				return
			} else if err == errorenum.GabisaAksesBro {
				r.log.Error(ctx, err.Error())
				c.JSON(http.StatusUnauthorized, payload.NewErrorResponse(err, traceID))
				return
			}
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		jsonRes.AccessToken = res.AccessToken

		r.log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
