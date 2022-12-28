package application

import (
	"vikishptra/domain_goarisan/controller/arisanapi"
	"vikishptra/domain_goarisan/gateway/withgorm"
	"vikishptra/domain_goarisan/usecase/rungruparisancreate"
	"vikishptra/domain_goarisan/usecase/runjoindetailgruparisan"
	"vikishptra/domain_goarisan/usecase/runusercreate"
	"vikishptra/domain_goarisan/usecase/runuserupdate"
	"vikishptra/shared/gogen"
	"vikishptra/shared/infrastructure/config"
	"vikishptra/shared/infrastructure/logger"
	"vikishptra/shared/infrastructure/server"
	"vikishptra/shared/infrastructure/token"
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

	jwtToken := token.NewJWTToken(cfg.JWTSecretKey)

	datasource := withgorm.NewGateway(log, appData, cfg)

	httpHandler := server.NewGinHTTPHandler(log, cfg.Servers[appName].Address, appData)

	x := arisanapi.NewGinController(log, cfg, jwtToken)
	x.AddUsecase(
		//
		runjoindetailgruparisan.NewUsecase(datasource),
		rungruparisancreate.NewUsecase(datasource),
		runuserupdate.NewUsecase(datasource),

		runusercreate.NewUsecase(datasource),
	)
	x.RegisterRouter(httpHandler.Router)

	httpHandler.RunWithGracefullyShutdown()

	return nil
}
