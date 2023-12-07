package genkube

import (
	"context"
	"testing"

	v1 "github.com/StevenCyb/genkube/test/api/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"go.uber.org/zap"
)

func TestFake(t *testing.T) {
	t.Parallel()

	logger := zap.NewNop().Sugar()
	fake, err := NewFake(logger, v1.AddToScheme)
	require.NoError(t, err)

	err = fake.Create(context.Background(), &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-pod",
			Namespace: "my-namespace",
		},
	})
	require.NoError(t, err)

	defer fake.Close()

	podList := &corev1.PodList{}
	err = fake.List(context.Background(), podList)
	require.NoError(t, err)
	require.NotNil(t, podList)
	assert.Len(t, podList.Items, 1)
	assert.Equal(t, "my-pod", podList.Items[0].Name)
	assert.Equal(t, "my-namespace", podList.Items[0].Namespace)
}
