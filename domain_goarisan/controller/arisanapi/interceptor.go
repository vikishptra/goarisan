package arisanapi

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"vikishptra/shared/infrastructure/token"
)

func (r *ginController) AuthMid() gin.HandlerFunc {

	return func(c *gin.Context) {
		//meriksa cookie token dan auth token
		var access_token string
		getAuth := token.ExtractToken(c)
		cookie, _ := c.Cookie("token")
		access_token = cookie
		if access_token != getAuth || access_token == "" {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

	}

}
