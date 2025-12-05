package main

import (
	"fmt"
	"net/http"
	"os"

	"lastlegends-gateway-service/internal/transport"

	"github.com/joho/godotenv"
	logger "github.com/ranggadablues/gosok/logger"
)

func main() {
	log := logger.NewLogger()

	// Load .env into environment
	if err := godotenv.Load(".env"); err != nil {
		log.LogErrorLevel("msg", "No .env file found: "+err.Error())
		return
	}

	handler := transport.NewHTTPHandler()

	gatewayPort := os.Getenv("GATEWAY_SERVICE_PORT")
	gatewayRunPort := fmt.Sprintf(":%s", gatewayPort)
	message := fmt.Sprintf("Gateway running on %s", gatewayRunPort)
	log.LogInfoLevel("msg", message)

	if err := http.ListenAndServe(gatewayRunPort, handler); err != nil {
		log.LogErrorLevel("msg", err.Error())
		panic(err)
	}
}
