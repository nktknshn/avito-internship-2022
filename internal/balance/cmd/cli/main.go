package main

import (
	_ "github.com/nktknshn/avito-internship-2022/internal/balance/adapters/cli/all"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/cli/root"
)

func main() {
	err := root.RootCmd.Execute()

	if err != nil {
		panic(err)
	}
}
