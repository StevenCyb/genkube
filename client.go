package genkube

import (
	"errors"

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ErrNoNoScheme is thrown when client is created without any scheme function.
var ErrNoNoScheme = errors.New("addToSchemeFuncs is empty, use e.g. corev1 or other crd-groups")

// AddToSchemeFunc define a function that appends a scheme.
type AddToSchemeFunc func(s *runtime.Scheme) error

// Client that wraps the common dynamic client and
// provide basic functionality of the Kubernetes API.
type Client struct {
	logger Logger
	close  chan struct{}
	client.WithWatch
}

// New creates a new client for multiple resource groups.
// Leave `kubeCredentialFile` empty to use cluster configuration.
// By default, no scheme is registered.
func New(logger Logger, kubeCredentialFile string, addToSchemeFuncs ...AddToSchemeFunc) (*Client, error) {
	if len(addToSchemeFuncs) == 0 {
		return nil, ErrNoNoScheme
	}

	var err error
	var restConfig *rest.Config
	if kubeCredentialFile != "" {
		apiConfig, err := clientcmd.LoadFromFile(kubeCredentialFile)
		if err != nil {
			return nil, err
		}

		restConfig, err = clientcmd.NewDefaultClientConfig(*apiConfig, &clientcmd.ConfigOverrides{}).ClientConfig()
		if err != nil {
			return nil, err
		}
	} else {
		if restConfig, err = rest.InClusterConfig(); err != nil {
			return nil, err
		}
	}

	scheme := runtime.NewScheme()
	for _, addToSchemeFunc := range addToSchemeFuncs {
		utilruntime.Must(addToSchemeFunc(scheme))
	}

	kubeClient, err := client.NewWithWatch(restConfig, client.Options{Scheme: scheme})
	if err != nil {
		return nil, err
	}

	return &Client{
		logger:    logger,
		WithWatch: kubeClient,
	}, nil
}

// Close calls the stop chang and stops all running watchers.
func (c *Client) Close() {
	close(c.close)
	c.close = make(chan struct{})
}
