# genkube
This is a simple wrapper client that inherit the `sigs.k8s.io/controller-runtime/pkg/client` client and adds some logic around it like watching any resource (including CRDs).

## Functions
The following functions are provided by the client. 
Most of them are part of the original client from *sigs.k8s.io* (including description). 

```go
  // Get retrieves an obj for the given object key from the Kubernetes Cluster.
  // obj must be a struct pointer so that obj can be updated with the response
  // returned by the Server.
  Get(ctx context.Context, key ObjectKey, obj Object, opts ...GetOption) error
```

```go
  // List retrieves list of objects for a given namespace and list options. On a
  // successful call, Items field in the list will be populated with the
  // result returned from the server.
  List(ctx context.Context, list ObjectList, opts ...ListOption) error
```

```go
  // Create saves the object obj in the Kubernetes cluster. obj must be a
  // struct pointer so that obj can be updated with the content returned by the Server.
  Create(ctx context.Context, obj Object, opts ...CreateOption) error
```

```go
  // Delete deletes the given obj from Kubernetes cluster.
  Delete(ctx context.Context, obj Object, opts ...DeleteOption) error
```

```go
  // Update updates the given obj in the Kubernetes cluster. obj must be a
  // struct pointer so that obj can be updated with the content returned by the Server.
  Update(ctx context.Context, obj Object, opts ...UpdateOption) error
```

```go
  // Patch patches the given obj in the Kubernetes cluster. obj must be a
  // struct pointer so that obj can be updated with the content returned by the Server.
  Patch(ctx context.Context, obj Object, patch Patch, opts ...PatchOption) error
```

```go
  // DeleteAllOf deletes all objects of the given type matching the given options.
  DeleteAllOf(ctx context.Context, obj Object, opts ...DeleteAllOfOption) error
```

```go
  // WatchResource watches a resource based on given objects lists.
  func (c *Client) WatchResource(
    resource runtime.Object, resourceList runtimeClient.ObjectList,
    eventHandler cache.ResourceEventHandlerFuncs, opts ...runtimeClient.ListOption,
  )
```