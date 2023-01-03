package arisanapi

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"vikishptra/domain_goarisan/model/errorenum"
	"vikishptra/domain_goarisan/usecase/runupdategruparisanbyidowner"
	"vikishptra/shared/gogen"
	"vikishptra/shared/infrastructure/logger"
	"vikishptra/shared/infrastructure/token"
	"vikishptra/shared/model/payload"
	"vikishptra/shared/util"
)

func (r *ginController) runupdategruparisanbyidownerHandler() gin.HandlerFunc {

	type InportRequest = runupdategruparisanbyidowner.InportRequest
	type InportResponse = runupdategruparisanbyidowner.InportResponse

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

		var jsonReqURI request
		if err := c.ShouldBindUri(&jsonReqURI); err != nil {
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

		var req InportRequest
		req.NamaGrup = jsonReqJSON.NamaGrup
		req.RulesMoney = jsonReqJSON.RulesMoney
		req.IDUser = jsonReqURI.IDUser
		req.IDGrup = jsonReqURI.IDGrup
		req.JwtToken = id

		r.log.Info(ctx, util.MustJSON(req))

		res, err := inport.Execute(ctx, req)
		if err != nil {
			if err == errorenum.DataNotFound {
				r.log.Error(ctx, err.Error())
				c.JSON(http.StatusNotFound, payload.NewErrorResponse(err, traceID))
				return
			} else if err == errorenum.HayoMauNgapain {
				r.log.Error(ctx, err.Error())
				c.JSON(http.StatusForbidden, payload.NewErrorResponse(err, traceID))
			} else if err == errorenum.AndaBukanAdmin {
				r.log.Error(ctx, err.Error())
				c.JSON(http.StatusForbidden, payload.NewErrorResponse(err, traceID))
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
