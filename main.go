package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/getsentry/sentry-go"

	"vikishptra/application"
	"vikishptra/shared/gogen"
)

var Version = "0.0.1"

func main() {
	appMap := map[string]gogen.Runner{
		//
		"apparisan": application.NewApparisan(),
	}

	flag.Parse()

	app, exist := appMap[flag.Arg(0)]
	if !exist {
		fmt.Println("You may try 'go run main.go <app_name>' :")
		for appName := range appMap {
			fmt.Printf(" - %s\n", appName)
		}
		return
	}

	fmt.Printf("Version %s\n", Version)
	err := app.Run()
	if err != nil {
		return
	}
	LogSentry()

}

func LogSentry() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://0bb83adb841a46fbbbdc54ecfb45d6b4@o4504520878718976.ingest.sentry.io/4504520880619520",
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)

	sentry.CaptureMessage("Hehhh")
}

// func openbrowser(url string) {
// 	var err error
//
// 	switch runtime.GOOS {
// 	case "linux":
// 		err = exec.Command("xdg-open", url).Start()
// 	case "windows":
// 		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
// 	case "darwin":
// 		err = exec.Command("open", url).Start()
// 	default:
// 		err = fmt.Errorf("unsupported platform")
// 	}
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// }
