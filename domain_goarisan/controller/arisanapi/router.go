package arisanapi

import (
	"github.com/gin-gonic/gin"

	"vikishptra/shared/gogen"
	"vikishptra/shared/infrastructure/config"
	"vikishptra/shared/infrastructure/logger"
	"vikishptra/shared/infrastructure/token"
)

type selectedRouter = gin.IRouter

type ginController struct {
	*gogen.BaseController
	log      logger.Logger
	cfg      *config.Config
	jwtToken token.JWTToken
}

func NewGinController(log logger.Logger, cfg *config.Config, tk token.JWTToken) gogen.RegisterRouterHandler[selectedRouter] {
	return &ginController{
		BaseController: gogen.NewBaseController(),
		log:            log,
		cfg:            cfg,
		jwtToken:       tk,
	}
}

func (r *ginController) RegisterRouter(router selectedRouter) {

	router.POST("/register", r.runUserCreateHandler())
	router.POST("/login", r.runUserLoginHandler())

	resource := router.Group("/api/v1", r.authentication())

	//fitur utama
	resource.PUT("/user/:id", r.authorization(), r.runUserUpdateHandler())
	resource.POST("/user/:id/grup", r.authorization(), r.runGrupArisanCreateHandler())
	resource.POST("/user/:id/join/grup", r.authorization(), r.runJoinDetailGrupArisanHandler())
	resource.POST("/user/:id/arisan/grup/", r.authorization(), r.runKocokGrupArisanHandler())
	resource.POST("/user/logout", r.authorization(), r.runLogoutUserHandler())
}
