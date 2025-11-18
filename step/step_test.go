package step

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bitrise-io/go-steputils/v2/stepconf"
	"github.com/bitrise-io/go-utils/v2/log"
	"github.com/bitrise-steplib/bitrise-step-gcp-storage-upload/step/mocks"
)

func TestConfigParsing(t *testing.T) {
	config := Config{
		Path:         "path",
		BucketName:   "bucket-name",
		BucketPrefix: "bucket-prefix",
		AccessToken:  "access-token",
	}

	mockEnvRepository := mocks.NewRepository(t)
	mockEnvRepository.On("Get", "path").Return(config.Path)
	mockEnvRepository.On("Get", "bucket_name").Return(config.BucketName)
	mockEnvRepository.On("Get", "bucket_prefix").Return(config.BucketPrefix)
	mockEnvRepository.On("Get", "access_token").Return(string(config.AccessToken))
	mockEnvRepository.On("Get", "verbose").Return("false")

	inputParser := stepconf.NewInputParser(mockEnvRepository)
	sut := NewUploader(inputParser, mockEnvRepository, log.NewLogger())

	receivedConfig, err := sut.ProcessConfig()
	assert.NoError(t, err)
	assert.Equal(t, config, receivedConfig)

	mockEnvRepository.AssertExpectations(t)
}
