package step

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"

	"golang.org/x/oauth2"

	"github.com/bitrise-io/go-steputils/v2/stepconf"
	"github.com/bitrise-io/go-utils/v2/env"
	"github.com/bitrise-io/go-utils/v2/log"
)

type Uploader struct {
	inputParser   stepconf.InputParser
	envRepository env.Repository
	logger        log.Logger
}

func NewUploader(inputParser stepconf.InputParser, envRepository env.Repository, logger log.Logger) Uploader {
	return Uploader{
		inputParser:   inputParser,
		envRepository: envRepository,
		logger:        logger,
	}
}

func (u Uploader) ProcessConfig() (Config, error) {
	var input Input
	err := u.inputParser.Parse(&input)
	if err != nil {
		return Config{}, err
	}

	stepconf.Print(input)
	u.logger.Println()
	u.logger.EnableDebugLog(input.Verbose)

	return Config{
		Path:         input.Path,
		BucketName:   input.BucketName,
		BucketPrefix: input.BucketPrefix,
		AccessToken:  input.AccessToken,
	}, nil
}

func (u Uploader) Run(config Config) error {
	items, err := u.collectItems(config.Path)
	if err != nil {
		return err
	}

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: string(config.AccessToken)})
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithTokenSource(ts))
	if err != nil {
		return fmt.Errorf("failed to create storage client: %w", err)
	}
	defer func(client *storage.Client) {
		if err := client.Close(); err != nil {
			u.logger.Warnf("Failed to close storage client: %s", err)
		}
	}(client)

	for _, item := range items {
		u.logger.Infof("Uploading file: %s", item.Path)

		f, err := os.Open(item.Path)
		if err != nil {
			return fmt.Errorf("open %s: %w", item.Path, err)
		}

		bucketPath := filepath.Join(config.BucketPrefix, item.Key)
		w := client.Bucket(config.BucketName).Object(bucketPath).NewWriter(ctx)
		w.ContentType = item.ContentType

		_, copyErr := io.Copy(w, f)
		_ = w.Close() // close writer to finalize the upload

		if closeErr := f.Close(); closeErr != nil {
			u.logger.Warnf("Failed to close file %s: %s", item.Path, closeErr)
		}

		if copyErr != nil {
			return fmt.Errorf("upload %s: %w", item.Path, err)
		}
	}

	return nil
}
