package arisanapi

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"vikishptra/domain_goarisan/model/errorenum"
	"vikishptra/domain_goarisan/model/vo"
	"vikishptra/domain_goarisan/usecase/runjoindetailgruparisan"
	"vikishptra/shared/gogen"
	"vikishptra/shared/infrastructure/logger"
	"vikishptra/shared/infrastructure/token"
	"vikishptra/shared/model/payload"
	"vikishptra/shared/util"
)

func (r *ginController) runJoinDetailGrupArisanHandler() gin.HandlerFunc {

	type InportRequest = runjoindetailgruparisan.InportRequest
	type InportResponse = runjoindetailgruparisan.InportResponse

	inport := gogen.GetInport[InportRequest, InportResponse](r.GetUsecase(InportRequest{}))

	type request struct {
		ReqIdGrup vo.GruparisanID `json:"id_detail_grup"`
		InportRequest
	}

	type response struct {
		InportResponse
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID()

		ctx := logger.SetTraceID(context.Background(), traceID)
		id, _ := token.ExtractTokenID(c)

		var jsonReqJSON request
		if err := c.BindJSON(&jsonReqJSON); err != nil {
			r.log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var req InportRequest
		req.ReqDetail.ID_Detail_Grup = jsonReqJSON.ReqIdGrup
		req.ReqDetail.ID_User = id
		req.ReqDetail.RandomString = util.GenerateID()
		req.ReqDetail.Now = time.Now()
		req.ReqDetail.JwtToken = id

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
