package arisanapi

import (
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"vikishptra/domain_goarisan/model/errorenum"
	"vikishptra/domain_goarisan/usecase/verifyemail"
	"vikishptra/shared/gogen"
	"vikishptra/shared/infrastructure/logger"
	"vikishptra/shared/model/payload"
	"vikishptra/shared/util"
)

func (r *ginController) verifyEmailHandler() gin.HandlerFunc {

	type InportRequest = verifyemail.InportRequest
	type InportResponse = verifyemail.InportResponse

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
			if err == errorenum.EmailSudahDiKonfirmasi {
				c.Redirect(http.StatusFound, os.Getenv("DOMAIN_REDIRECT")+"/emailconfirm")
			} else if err == errorenum.KonfirmasiEmailAndaSudahKadaluawarsa {
				c.Redirect(http.StatusFound, os.Getenv("DOMAIN_REDIRECT")+"/emailkadaluarsa")
			}
			r.log.Error(ctx, err.Error())
			c.Redirect(http.StatusFound, os.Getenv("DOMAIN_REDIRECT")+"/kesalahanuser")
			return
		}

		var jsonRes response
		jsonRes.Message = res.Message

		r.log.Info(ctx, util.MustJSON(jsonRes))
		c.Redirect(http.StatusFound, os.Getenv("DOMAIN_REDIRECT")+"/emailsuccess")
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
