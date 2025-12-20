package main

import (
	"fmt"
	"net/http"
	"os"

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

	if err := http.ListenAndServe(gatewayRunPort, handler); err != nil {
		log.LogErrorLevel("msg", err.Error())
		panic(err)
	}
}
