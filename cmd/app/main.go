package main

import (
	"github.com/varkis-ms/service-api-gateway/internal/app"
)

const configsDir = "."

func main() {
	app.Run(configsDir)
}
