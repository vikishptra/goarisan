package arisanapi

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"vikishptra/domain_goarisan/usecase/rungruparisancreate"
	"vikishptra/shared/gogen"
	"vikishptra/shared/infrastructure/logger"
	"vikishptra/shared/model/payload"
	"vikishptra/shared/util"
)

func (r *ginController) runGrupArisanCreateHandler() gin.HandlerFunc {

	type InportRequest = rungruparisancreate.InportRequest
	type InportResponse = rungruparisancreate.InportResponse

	inport := gogen.GetInport[InportRequest, InportResponse](r.GetUsecase(InportRequest{}))

	type request struct {
		InportRequest
	}

	type response struct {
		Message string `json:"message"`
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID()

		ctx := logger.SetTraceID(context.Background(), traceID)

		var jsonReqUri request
		if err := c.BindUri(&jsonReqUri); err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}
		var jsonReq request
		if err := c.BindJSON(&jsonReq); err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}
		var req InportRequest
		req.ID_Owner = jsonReqUri.ID_Owner
		req.NamaGrup = jsonReq.NamaGrup
		req.JumlahUsers = jsonReq.JumlahUsers
		req.ResultMoney = jsonReq.RulesMoney
		req.RulesMoney = jsonReq.RulesMoney
		req.Now = time.Now()
		req.RandomString = util.GenerateID()
		req.DetailReq.RandomString = util.GenerateID()

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
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
