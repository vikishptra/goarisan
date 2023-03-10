package arisanapi

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"vikishptra/domain_goarisan/model/errorenum"
	"vikishptra/domain_goarisan/usecase/runkocokgruparisan"
	"vikishptra/shared/gogen"
	"vikishptra/shared/infrastructure/logger"
	"vikishptra/shared/infrastructure/token"
	"vikishptra/shared/model/payload"
	"vikishptra/shared/util"
)

func (r *ginController) runKocokGrupArisanHandler() gin.HandlerFunc {

	type InportRequest = runkocokgruparisan.InportRequest
	type InportResponse = runkocokgruparisan.InportResponse

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
		if err := c.BindUri(&jsonReqURI); err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var req InportRequest

		req.IDGrup = jsonReqURI.IDGrup
		req.IDUser = id
		req.JwtToken = id
		r.log.Info(ctx, util.MustJSON(req))

		res, err := inport.Execute(ctx, req)
		if err != nil {
			if err == errorenum.AndaBukanAdmin || err == errorenum.HayoMauNgapain {
				r.log.Error(ctx, err.Error())
				c.JSON(http.StatusForbidden, payload.NewErrorResponse(err, traceID))
				return
			} else if err == errorenum.SomethingError {
				r.log.Error(ctx, err.Error())
				c.JSON(http.StatusInternalServerError, payload.NewErrorResponse(err, traceID))
				return
			} else if err == errorenum.DataNotFound {
				r.log.Error(ctx, err.Error())
				c.JSON(http.StatusNotFound, payload.NewErrorResponse(err, traceID))
				return
			}
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		jsonRes.Items = res.Items

		r.log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
