package main

import (
	"context"

	"github.com/hserge/namak/app"
	_ "github.com/hserge/namak/docs"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8888
// @BasePath /
func main() {
	ctx := context.Background()
	app.Initialize(ctx)
}
