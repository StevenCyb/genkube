package genkube

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
	runtimeClient "sigs.k8s.io/controller-runtime/pkg/client"
)

// WatchResource watches a resource based on given objects lists.
func (c *Client) WatchResource(
	resource runtime.Object, resourceList runtimeClient.ObjectList,
	eventHandler cache.ResourceEventHandlerFuncs, opts ...runtimeClient.ListOption,
) {
	_, controller := cache.NewInformer(
		&cache.ListWatch{
			ListFunc: func(lo metav1.ListOptions) (result runtime.Object, err error) {
				err = c.List(context.Background(), resourceList, opts...)
				return resourceList, err
			},
			WatchFunc: func(lo metav1.ListOptions) (watch.Interface, error) {
				lo.Watch = true
				watcher, err := c.Watch(
					context.Background(),
					resourceList,
					opts...,
				)

				return watcher, err
			},
		}, resource, 0, eventHandler,
	)

	go controller.Run(c.close)
}
