package main

import (
	"os"

	"github.com/bitrise-io/go-steputils/v2/stepconf"
	"github.com/bitrise-io/go-steputils/v2/stepenv"
	"github.com/bitrise-io/go-utils/v2/env"
	. "github.com/bitrise-io/go-utils/v2/exitcode"
	"github.com/bitrise-io/go-utils/v2/log"
	"github.com/bitrise-steplib/bitrise-step-gcp-storage-upload/step"
)

func main() {
	exitCode := run()
	os.Exit(int(exitCode))
}

func run() ExitCode {
	logger := log.NewLogger()

	uploader := createUploader(logger)
	config, err := uploader.ProcessConfig()
	if err != nil {
		logger.Errorf("Process config: %s", err)
		return Failure
	}

	if err := uploader.Run(config); err != nil {
		logger.Errorf("Run: %s", err)
		return Failure
	}

	return Success
}

func createUploader(logger log.Logger) step.Uploader {
	envRepository := stepenv.NewRepository(env.NewRepository())
	inputParser := stepconf.NewInputParser(envRepository)

	return step.NewUploader(inputParser, envRepository, logger)
}
