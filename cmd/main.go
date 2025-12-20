package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"lastlegends-gateway-service/internal/transport"

	appEmbed "lastlegends-gateway-service"

	"github.com/joho/godotenv"
	logger "github.com/ranggadablues/gosok/logger"
)

func main() {
	log := logger.NewLogger()

	// Load .env into environment
	//
	dir, err := os.Getwd()
	if err != nil {
		log.LogErrorLevel("msg", "Dir not found with error: "+err.Error())
	}
	env := fmt.Sprintf("%s/.env", dir)
	if err := godotenv.Load(env); err != nil {
		log.LogErrorLevel("msg", "No .env file found: "+err.Error())
	}

	handler := transport.NewHTTPHandler(
		transport.SwaggerAssets{
			JSON: appEmbed.SwaggerJSON,
			UI:   appEmbed.SwaggerUI,
		},
	)

	gatewayPort := os.Getenv("GATEWAY_SERVICE_PORT")
	gatewayRunPort := fmt.Sprintf(":%s", gatewayPort)
	message := fmt.Sprintf("Gateway running on %s", gatewayRunPort)
	log.LogInfoLevel("msg", message)

	srv := &http.Server{
		Addr:         gatewayRunPort,
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.LogErrorLevel("msg", err.Error())
		panic(err)
	}
}
