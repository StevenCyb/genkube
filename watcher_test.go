package genkube

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	fakeclient "sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestWatchResource(t *testing.T) {
	t.Parallel()

	fakeClient := fakeclient.NewClientBuilder().Build()
	testClient := &Client{
		WithWatch: fakeClient,
		close:     make(chan struct{}),
	}
	fakePod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-pod",
			Namespace: "my-namespace",
		},
	}

	fakePodList := &corev1.PodList{}
	eventHandler := cache.ResourceEventHandlerFuncs{
		AddFunc:    func(obj interface{}) {},
		UpdateFunc: func(oldObj, newObj interface{}) {},
		DeleteFunc: func(obj interface{}) {},
	}

	stopCh := make(chan struct{})
	defer close(stopCh)

	go testClient.WatchResource(fakePod, fakePodList, eventHandler, client.InNamespace("my-namespace"))

	err := fakeClient.Create(context.Background(), fakePod)
	require.NoError(t, err)

	time.Sleep(time.Millisecond * 100)
	testClient.Close()
}
