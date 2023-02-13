package application

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"

	"vikishptra/domain_goarisan/controller/arisanapi"
	"vikishptra/domain_goarisan/gateway/withgorm"
	"vikishptra/domain_goarisan/usecase/deletedetailgrupbyowner"
	"vikishptra/domain_goarisan/usecase/findgruparisanbyidowner"
	"vikishptra/domain_goarisan/usecase/findgrupbyiduser"
	"vikishptra/domain_goarisan/usecase/findoneuserbyid"
	"vikishptra/domain_goarisan/usecase/refreshtokenjwt"
	"vikishptra/domain_goarisan/usecase/runchangepasswordwithgmail"
	"vikishptra/domain_goarisan/usecase/runcreatepaymentmidtrans"
	"vikishptra/domain_goarisan/usecase/rungruparisancreate"
	"vikishptra/domain_goarisan/usecase/runjoindetailgruparisan"
	"vikishptra/domain_goarisan/usecase/runkocokgruparisan"
	"vikishptra/domain_goarisan/usecase/runlogoutuser"
	"vikishptra/domain_goarisan/usecase/runnewpasswordconfirmemail"
	"vikishptra/domain_goarisan/usecase/runnotificationmidtrans"
	"vikishptra/domain_goarisan/usecase/runupdatdetailgruparisans"
	"vikishptra/domain_goarisan/usecase/runupdategruparisanbyidowner"
	"vikishptra/domain_goarisan/usecase/runupdateownergrup"
	"vikishptra/domain_goarisan/usecase/runupdateusermoney"
	"vikishptra/domain_goarisan/usecase/runusercreate"
	"vikishptra/domain_goarisan/usecase/runuserlogin"
	"vikishptra/domain_goarisan/usecase/runuserupdate"
	"vikishptra/domain_goarisan/usecase/runverifynewpasswordemail"
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
	LogSentry()

	x := arisanapi.NewGinController(log, cfg)
	_, err := os.LookupEnv("PORT")
	x.AddUsecase(
		//
		runnotificationmidtrans.NewUsecase(datasource),
		runcreatepaymentmidtrans.NewUsecase(datasource),
		runnewpasswordconfirmemail.NewUsecase(datasource),
		runverifynewpasswordemail.NewUsecase(datasource),
		runchangepasswordwithgmail.NewUsecase(datasource),
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
	if err {
		httpHandler.Router.Run()
	}
	httpHandler.Router.Use(sentrygin.New(sentrygin.Options{}))

	httpHandler.RunWithGracefullyShutdown()

	return nil
}
func LogSentry() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              "https://0bb83adb841a46fbbbdc54ecfb45d6b4@o4504520878718976.ingest.sentry.io/4504520880619520",
		EnableTracing:    true,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)
	defer func() {
		if r := recover(); r != nil {
			sentry.CaptureException(fmt.Errorf("%v", r))
			sentry.Flush(2 * time.Second)
		}
	}()
	if err != nil {
		sentry.CaptureException(err)
	}

	sentry.CaptureMessage("error bro")
}
