//nolint:gochecknoglobals
package v1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

const (
	// GroupVersionGroup specifies the group.
	GroupVersionGroup = "post.funny.com"
	// GroupVersionVersion specifies the version.
	GroupVersionVersion = "v1"
	// APIVersion specifies the api version string.
	APIVersion = GroupVersionGroup + "/" + GroupVersionVersion
)

var (
	// GroupVersion is group version used to register these objects.
	GroupVersion = schema.GroupVersion{Group: GroupVersionGroup, Version: GroupVersionVersion}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme.
	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}

	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme
)
