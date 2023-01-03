package main

import (
	"context"

	"github.com/hserge/namak/app"
)

func main() {
	ctx := context.Background()
	app.Initialize(ctx)
}
