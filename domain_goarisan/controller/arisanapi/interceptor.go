package arisanapi

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"vikishptra/domain_goarisan/model/errorenum"
	"vikishptra/shared/infrastructure/token"
	"vikishptra/shared/model/payload"
	"vikishptra/shared/util"
)

func (r *ginController) AuthMid() gin.HandlerFunc {

	return func(c *gin.Context) {
		traceID := util.GenerateID()
		//meriksa cookie token dan auth token
		var access_token string
		getAuth := token.ExtractToken(c)
		cookie, _ := c.Cookie("token")
		err := errorenum.GabisaAksesBro

		access_token = cookie
		if access_token != getAuth || access_token == "" {
			c.JSON(http.StatusUnauthorized, payload.NewErrorResponse(err, traceID))
			c.Abort()
			return
		}

	}

}
