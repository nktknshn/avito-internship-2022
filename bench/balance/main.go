package main

import (
	"os"

	_ "github.com/nktknshn/avito-internship-2022-bench/cli/all"
	"github.com/nktknshn/avito-internship-2022-bench/cli/root"
	"github.com/nktknshn/avito-internship-2022-bench/logger"
)

func main() {
	log := logger.GetLogger()

	log.Info("Starting balance benchmark")
	if err := root.RootCmd.Execute(); err != nil {
		log.Error("Failed to execute balance benchmark", "error", err)
		os.Exit(1)
	}
}
