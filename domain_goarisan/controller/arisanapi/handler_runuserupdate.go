package arisanapi

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"vikishptra/domain_goarisan/controller/arisanapi/token"
	"vikishptra/domain_goarisan/model/errorenum"
	"vikishptra/domain_goarisan/usecase/runuserupdate"
	"vikishptra/shared/gogen"
	"vikishptra/shared/infrastructure/logger"
	"vikishptra/shared/model/payload"
	"vikishptra/shared/util"
)

func (r *ginController) runUserUpdateHandler() gin.HandlerFunc {

	type InportRequest = runuserupdate.InportRequest
	type InportResponse = runuserupdate.InportResponse

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
		id, _ := token.ExtractTokenID(c)

		if err := c.BindUri(&jsonReq); err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}
		var jsonReqJSON request
		if err := c.BindJSON(&jsonReqJSON); err != nil {
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

		var req InportRequest
		req.ID = jsonReq.ID
		req.Jwt = id
		req.Name = jsonReqJSON.Name

		r.log.Info(ctx, util.MustJSON(req))

		res, err := inport.Execute(ctx, req)
		if err != nil {
			r.log.Error(ctx, err.Error())
			if err == errorenum.DataNotFound {
				c.JSON(http.StatusNotFound, payload.NewErrorResponse(err, traceID))
				return
			} else if err == errorenum.HayoMauNgapain {
				c.JSON(http.StatusForbidden, payload.NewErrorResponse(err, traceID))
				return
			}
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		jsonRes.ID = res.ID
		jsonRes.Nama = res.Nama
		jsonRes.Message = res.Message
		r.log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
