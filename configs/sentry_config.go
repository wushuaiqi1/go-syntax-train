package configs

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

func LoadSentryConfig() {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              viper.GetString("sentry.dsn"),
		TracesSampleRate: 1.0,
		ServerName:       viper.GetString("server.name"),
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	// Create an instance of sentryhttp
	sentryHandler = sentryhttp.New(sentryhttp.Options{
		WaitForDelivery: true,
		Timeout:         time.Second,
	})

	// Once it's done, you can set up routes and attach the handler as one of your middleware
	log.Println("load sentry config success...")
}

var sentryHandler *sentryhttp.Handler

func RegisterRouteBindSentry(pattern string, function http.HandlerFunc) {
	http.HandleFunc(pattern, sentryHandler.HandleFunc(function))
}
