package arisanapi

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"vikishptra/domain_goarisan/controller/arisanapi/token"
)

func (r *ginController) authentication() gin.HandlerFunc {
	return func(c *gin.Context) {

		// tokenInBytes, err := r.JwtToken.VerifyToken(c.GetHeader("token"))
		// if err != nil {
		// 	c.AbortWithStatus(http.StatusForbidden)
		// 	return
		// }
		//
		// var dataToken payload.DataToken
		// err = json.Unmarshal(tokenInBytes, &dataToken)
		// if err != nil {
		// 	c.AbortWithStatus(http.StatusForbidden)
		// 	return
		// }
		//
		// c.Set("data", dataToken)
		//
		// c.AbortWithStatus(http.StatusForbidden)
		// return

	}
}

func (r *ginController) authorization() gin.HandlerFunc {

	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
