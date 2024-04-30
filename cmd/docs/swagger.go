package main

import (
	"os"

	log "github.com/robowealth-mutual-fund/stdlog"

	"github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/internals/config"
	"github.com/robowealth-mutual-fund/blueprint-roa-fiber-golang/openapi"
)

func main() {
	swagger := openapi.NewOpenAPI3(config.New())
	swaggerJSONBytes, err := swagger.MarshalJSON()

	if err != nil {
		log.Error("Failed to marshal swagger json", err)

		return
	}

	if err = os.WriteFile("./www/swagger-ui/swagger.json", swaggerJSONBytes, 0644); err != nil {
		log.Error("Failed to create swagger json file", err)
	}
}
