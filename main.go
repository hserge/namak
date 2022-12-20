package main

import (
	"context"

	"github.com/hserge/namak/app"
)

func main() {
	ctx := context.Background()

	a := app.App{}
	a.Initialize(ctx, "postgres://dbuser:password@localhost:5432/db")
	defer a.CloseDb()

	a.Start(":8888")
}
