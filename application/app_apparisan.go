package application

import (
	"os"

	"github.com/gin-contrib/cors"

	"vikishptra/domain_goarisan/controller/arisanapi"
	"vikishptra/domain_goarisan/gateway/withgorm"
	"vikishptra/domain_goarisan/usecase/deletedetailgrupbyowner"
	"vikishptra/domain_goarisan/usecase/findgruparisanbyidowner"
	"vikishptra/domain_goarisan/usecase/findgrupbyiduser"
	"vikishptra/domain_goarisan/usecase/findoneuserbyid"
	"vikishptra/domain_goarisan/usecase/refreshtokenjwt"
	"vikishptra/domain_goarisan/usecase/rungruparisancreate"
	"vikishptra/domain_goarisan/usecase/runjoindetailgruparisan"
	"vikishptra/domain_goarisan/usecase/runkocokgruparisan"
	"vikishptra/domain_goarisan/usecase/runlogoutuser"
	"vikishptra/domain_goarisan/usecase/runupdatdetailgruparisans"
	"vikishptra/domain_goarisan/usecase/runupdategruparisanbyidowner"
	"vikishptra/domain_goarisan/usecase/runupdateownergrup"
	"vikishptra/domain_goarisan/usecase/runupdateusermoney"
	"vikishptra/domain_goarisan/usecase/runusercreate"
	"vikishptra/domain_goarisan/usecase/runuserlogin"
	"vikishptra/domain_goarisan/usecase/runuserupdate"
	"vikishptra/domain_goarisan/usecase/sendemailconfirm"
	"vikishptra/domain_goarisan/usecase/verifyemail"
	"vikishptra/shared/gogen"
	"vikishptra/shared/infrastructure/config"
	"vikishptra/shared/infrastructure/logger"
	"vikishptra/shared/infrastructure/server"
)

type apparisan struct{}

func NewApparisan() gogen.Runner {
	return &apparisan{}
}

func (apparisan) Run() error {

	const appName = "apparisan"

	cfg := config.ReadConfig()

	appData := gogen.NewApplicationData(appName)

	log := logger.NewSimpleJSONLogger(appData)

	datasource := withgorm.NewGateway(log, appData, cfg)

	httpHandler := server.NewGinHTTPHandler(log, cfg.Servers[appName].Address, appData)

	x := arisanapi.NewGinController(log, cfg)
	_, err := os.LookupEnv("PORT")

	x.AddUsecase(
		//
		sendemailconfirm.NewUsecase(datasource),
		verifyemail.NewUsecase(datasource),
		refreshtokenjwt.NewUsecase(datasource),
		runupdateownergrup.NewUsecase(datasource),
		deletedetailgrupbyowner.NewUsecase(datasource),
		findgruparisanbyidowner.NewUsecase(datasource),
		runupdatdetailgruparisans.NewUsecase(datasource),
		findgrupbyiduser.NewUsecase(datasource),
		runupdategruparisanbyidowner.NewUsecase(datasource),
		runupdateusermoney.NewUsecase(datasource),
		findoneuserbyid.NewUsecase(datasource),
		runlogoutuser.NewUsecase(datasource),
		runuserlogin.NewUsecase(datasource),
		runkocokgruparisan.NewUsecase(datasource),
		runjoindetailgruparisan.NewUsecase(datasource),
		rungruparisancreate.NewUsecase(datasource),
		runuserupdate.NewUsecase(datasource),
		runusercreate.NewUsecase(datasource),
	)
	x.RegisterRouter(httpHandler.Router)

	corsConfig := cors.DefaultConfig()
	Origin := os.Getenv("ORIGIN")
	OriginUrl := Origin
	corsConfig.AllowOrigins = []string{OriginUrl}
	corsConfig.AllowCredentials = true

	if err {
		httpHandler.Router.Run()
	}
	httpHandler.Router.Use(cors.New(corsConfig))
	httpHandler.RunWithGracefullyShutdown()

	return nil
}
