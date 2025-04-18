package main

import "github.com/nktknshn/avito-internship-2022/internal/app"

func main() {
	app, err := app.NewApp()

	if err != nil {
		panic(err)
	}

	// cfg

	_ = app
}
