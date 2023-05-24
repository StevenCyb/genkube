package genkube

import (
	"testing"

	v1 "github.com/StevenCyb/genkube/test/api/v1"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestNewClient(t *testing.T) {
	t.Parallel()

	logger := zap.NewNop().Sugar()
	kubeCredentialFile := "test/kubeconfig"
	addToSchemeFuncs := []AddToSchemeFunc{
		v1.AddToScheme,
	}
	client, err := New(logger, kubeCredentialFile, addToSchemeFuncs...)

	assert.NoError(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, logger, client.logger)
	assert.NotNil(t, client.WithWatch)
}
