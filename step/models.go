package step

import (
	"github.com/bitrise-io/go-steputils/v2/stepconf"
)

type Input struct {
	Path         string          `env:"path,required"`
	BucketName   string          `env:"bucket_name,required"`
	BucketPrefix string          `env:"bucket_prefix"`
	AccessToken  stepconf.Secret `env:"access_token"`
	Verbose      bool            `env:"verbose,opt[true,false]"`
}

type Config struct {
	Path         string
	BucketName   string
	BucketPrefix string
	AccessToken  stepconf.Secret
}
