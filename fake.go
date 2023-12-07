package genkube

import (
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	fakeClient "sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func NewFake(logger *zap.SugaredLogger, addToSchemeFuncs ...AddToSchemeFunc) (*Client, error) {
	if len(addToSchemeFuncs) == 0 {
		return nil, ErrNoNoScheme
	}

	fc := fakeClient.NewClientBuilder().Build()

	scheme := runtime.NewScheme()
	for _, addToSchemeFunc := range addToSchemeFuncs {
		utilruntime.Must(addToSchemeFunc(scheme))
	}

	return &Client{
		logger:    logger,
		WithWatch: fc,
	}, nil
}
