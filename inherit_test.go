package genkube

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	fakeclient "sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestGet(t *testing.T) {
	t.Parallel()

	fakeClient := fakeclient.NewClientBuilder().
		WithObjects(&corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "my-pod",
				Namespace: "my-namespace",
			},
		}).
		Build()

	testClient := &Client{
		WithWatch: fakeClient,
	}

	obj := &corev1.Pod{}
	err := testClient.Get(context.TODO(), types.NamespacedName{
		Namespace: "my-namespace",
		Name:      "my-pod",
	}, obj)

	assert.NoError(t, err)
	assert.Equal(t, "my-pod", obj.Name)
}

func TestList(t *testing.T) {
	t.Parallel()

	fakeClient := fakeclient.NewClientBuilder().
		WithObjects(&corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "my-pod",
				Namespace: "my-namespace",
			},
		}).
		WithObjects(&corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "my-pod-2",
				Namespace: "my-namespace",
			},
		}).
		WithObjects(&corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "my-pod-3",
				Namespace: "other-namespace",
			},
		}).
		Build()

	testClient := &Client{
		WithWatch: fakeClient,
	}

	objs := &corev1.PodList{}
	err := testClient.List(context.TODO(), objs, client.InNamespace("my-namespace"))

	assert.NoError(t, err)
	assert.Len(t, objs.Items, 2)
}

func TestCreate(t *testing.T) {
	t.Parallel()

	fakeClient := fakeclient.NewClientBuilder().Build()
	testClient := &Client{
		WithWatch: fakeClient,
	}

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-pod",
			Namespace: "my-namespace",
		},
	}

	err := testClient.Create(context.TODO(), pod)

	assert.NoError(t, err)

	obj := &corev1.Pod{}
	err = fakeClient.Get(context.TODO(), types.NamespacedName{
		Namespace: "my-namespace",
		Name:      "my-pod",
	}, obj)

	assert.NoError(t, err)
	assert.Equal(t, pod.Name, obj.Name)
}

func TestDelete(t *testing.T) {
	t.Parallel()

	fakeClient := fakeclient.NewClientBuilder().Build()
	testClient := &Client{
		WithWatch: fakeClient,
	}

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-pod",
			Namespace: "my-namespace",
		},
	}

	// Create the pod first
	err := fakeClient.Create(context.TODO(), pod)
	assert.NoError(t, err)

	// Delete the pod
	err = testClient.Delete(context.TODO(), pod)
	assert.NoError(t, err)

	// Verify that the pod is deleted
	obj := &corev1.Pod{}
	err = fakeClient.Get(context.TODO(), types.NamespacedName{
		Namespace: "my-namespace",
		Name:      "my-pod",
	}, obj)
	assert.Error(t, err)
	assert.NoError(t, client.IgnoreNotFound(err))
}

func TestDeleteAllOf(t *testing.T) {
	t.Parallel()

	fakeClient := fakeclient.NewClientBuilder().Build()
	testClient := &Client{
		WithWatch: fakeClient,
	}

	pod1 := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-pod-1",
			Namespace: "my-namespace",
		},
	}

	pod2 := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-pod-2",
			Namespace: "my-namespace",
		},
	}

	// Create the pods first
	err := fakeClient.Create(context.TODO(), pod1)
	assert.NoError(t, err)
	err = fakeClient.Create(context.TODO(), pod2)
	assert.NoError(t, err)

	// Delete all pods in the namespace
	err = testClient.DeleteAllOf(context.TODO(), &corev1.Pod{}, client.InNamespace("my-namespace"))
	assert.NoError(t, err)

	// Verify that all pods are deleted
	pods := &corev1.PodList{}
	err = fakeClient.List(context.TODO(), pods, client.InNamespace("my-namespace"))
	assert.NoError(t, err)
	assert.Len(t, pods.Items, 0)
}

func TestPatch(t *testing.T) {
	t.Parallel()

	fakeClient := fakeclient.NewClientBuilder().Build()
	testClient := &Client{
		WithWatch: fakeClient,
	}

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-pod",
			Namespace: "my-namespace",
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "my-container",
					Image: "nginx:latest",
				},
			},
		},
	}

	// Create the pod first
	err := fakeClient.Create(context.TODO(), pod)
	assert.NoError(t, err)

	// Modify the pod's image using Patch
	modifiedPod := pod.DeepCopy()
	modifiedPod.Spec.Containers[0].Image = "nginx:modified"
	err = testClient.Patch(context.TODO(), modifiedPod, client.MergeFrom(pod.DeepCopy()))
	assert.NoError(t, err)

	// Retrieve the patched pod
	retrievedPod := &corev1.Pod{}
	err = fakeClient.Get(context.TODO(), types.NamespacedName{
		Namespace: "my-namespace",
		Name:      "my-pod",
	}, retrievedPod)
	assert.NoError(t, err)

	// Verify that the image has been modified
	assert.Equal(t, "nginx:modified", retrievedPod.Spec.Containers[0].Image)
}

func TestUpdate(t *testing.T) {
	t.Parallel()

	fakeClient := fakeclient.NewClientBuilder().Build()
	testClient := &Client{
		WithWatch: fakeClient,
	}

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-pod",
			Namespace: "my-namespace",
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "my-container",
					Image: "nginx:latest",
				},
			},
		},
	}

	// Create the pod first
	err := fakeClient.Create(context.TODO(), pod)
	assert.NoError(t, err)

	// Modify the pod's image
	pod.Spec.Containers[0].Image = "nginx:modified"

	// Update the pod
	err = testClient.Update(context.TODO(), pod)
	assert.NoError(t, err)

	// Retrieve the updated pod
	retrievedPod := &corev1.Pod{}
	err = fakeClient.Get(context.TODO(), types.NamespacedName{
		Namespace: "my-namespace",
		Name:      "my-pod",
	}, retrievedPod)
	assert.NoError(t, err)

	// Verify that the image has been updated
	assert.Equal(t, "nginx:modified", retrievedPod.Spec.Containers[0].Image)
}
