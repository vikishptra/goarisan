package arisanapi

import (
	"github.com/gin-gonic/gin"

	"vikishptra/shared/gogen"
	"vikishptra/shared/infrastructure/config"
	"vikishptra/shared/infrastructure/logger"
)

type selectedRouter = gin.IRouter

type ginController struct {
	*gogen.BaseController
	log logger.Logger
	cfg *config.Config
}

func NewGinController(log logger.Logger, cfg *config.Config) gogen.RegisterRouterHandler[selectedRouter] {
	return &ginController{
		BaseController: gogen.NewBaseController(),
		log:            log,
		cfg:            cfg,
	}
}

func (r *ginController) RegisterRouter(router selectedRouter) {

	router.POST("/register", r.runUserCreateHandler())
	router.POST("/login", r.runUserLoginHandler())

	resource := router.Group("/api/v1", r.AuthMid())

	//fitur utama
	resource.PUT("/user/:id", r.runUserUpdateHandler())
	resource.POST("/user/:id/grup", r.runGrupArisanCreateHandler())
	resource.POST("/user/:id/join/grup", r.runJoinDetailGrupArisanHandler())
	resource.POST("/user/:id/arisan/grup/", r.runKocokGrupArisanHandler())
	resource.POST("/user/logout", r.runLogoutUserHandler())
}
