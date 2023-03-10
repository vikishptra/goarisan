package arisanapi

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"

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
	bucket := ratelimit.NewBucket(time.Minute, 25)
	//umum
	router.POST("/register", r.runUserCreateHandler())
	router.POST("/login", r.runUserLoginHandler())
	router.GET("/refresh-token", r.refreshtokenjwtHandler())
	router.POST("/logout", r.runLogoutUserHandler())
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Vicky Sahputra GO-ARISAN"})
	})

	//email
	router.GET("/verifyemail", r.verifyEmailHandler())
	router.POST("/confirm", RateLimitMiddleware(bucket), r.sendemailconfirmHandler())

	//password
	router.POST("/change/password", RateLimitMiddleware(bucket), r.runChangePasswordWithGmailHandler())
	router.GET("/change/password", r.runverifynewpasswordemailHandler())
	router.POST("/new/password", r.runnewpasswordconfirmemailHandler())

	//notification push
	router.POST("/notification/push", r.runnotificationmidtransHandler())

	resource := router.Group("/api/v1", r.AuthMid())

	//fitur utama
	//user
	resource.PUT("/user/", r.runUserUpdateHandler())
	resource.GET("/user/", r.findOneUserByIDHandler())

	//grup
	resource.POST("/user/grup", r.runGrupArisanCreateHandler())
	resource.POST("/user/join-grup", r.runJoinDetailGrupArisanHandler())
	resource.GET("/user/owner/grup", r.findgruparisanbyidOwnerHandler())
	resource.PUT("/user/owner/:grup", r.runupdategruparisanbyidownerHandler())

	//detail_grup
	resource.GET("/user/grup", r.findgrupbyiduserHandler())
	resource.POST("/user/arisan/:grup", r.runKocokGrupArisanHandler())
	resource.PUT("/user/setor-arisan/:grup", r.runupdatdetailgruparisansHandler())
	resource.DELETE("/user/owner/:id_grup/:id_user", r.deletedetailgrupbyownerHandler())
	resource.POST("/user/owner", r.runUpdateOwnerGrupHandler())

	//payment
	resource.POST("/payment/bank-transfer", r.runcreatepaymentmidtransHandler())
}
