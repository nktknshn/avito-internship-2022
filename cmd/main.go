package main

import (
	"context"

	"github.com/nktknshn/avito-internship-2022/internal/config"
)

func main() {
	ctx := context.Background()
	app, err := NewApplication(ctx, &config.Config{})

	if err != nil {
		panic(err)
	}

	// cfg

	_ = app
}
