package arisanapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"

	"vikishptra/domain_goarisan/model/errorenum"
	"vikishptra/shared/infrastructure/token"
	"vikishptra/shared/model/payload"
	"vikishptra/shared/util"
)

func (r *ginController) AuthMid() gin.HandlerFunc {

	return func(c *gin.Context) {
		traceID := util.GenerateID()
		//meriksa cookie token dan auth token
		refreshToken := token.ExtractTokenCookie(c)
		if err := token.TokenValid(c); err != nil {
			c.JSON(http.StatusUnauthorized, payload.NewErrorResponse(errorenum.TokenExpired, traceID))
			c.Abort()
			return
		}
		if refreshToken == "" {
			c.JSON(http.StatusUnauthorized, payload.NewErrorResponse(errorenum.TokenExpired, traceID))
			c.Abort()
			return
		}
	}

}

func RateLimitMiddleware(bucket *ratelimit.Bucket) gin.HandlerFunc {
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) == 0 {
			traceID := util.GenerateID()
			c.JSON(http.StatusTooManyRequests, payload.NewErrorResponse(errorenum.TooManyRequests, traceID))
			c.Abort()
			return
		}
		c.Next()
	}
}
