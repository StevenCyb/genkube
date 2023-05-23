//nolint:gochecknoglobals
package v1

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	// KindPost is the kind for this resources.
	KindPost = "Post"
	// ResourceNetwork specifies the resource name of this resource.
	ResourcePost = "posts"
)

// PostActionsGroupVersionResource specifies the GVR for the resource.
var PostActionsGroupVersionResource = schema.GroupVersionResource{
	Group:    GroupVersion.Group,
	Version:  GroupVersion.Version,
	Resource: ResourcePost,
}

// Post is the Schema for the posts API
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Post struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"` // nolint:tagliatelle

	Spec   PostSpec   `json:"spec,omitempty"`
	Status PostStatus `json:"status,omitempty"`
}

// PostList contains a list of Post.
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type PostList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"` // nolint:tagliatelle
	Items           []Post                      `json:"items"`
}

func init() { // nolint:gochecknoinits
	SchemeBuilder.Register(&Post{}, &PostList{})
}

// PostSpec holds specifications for the resource.
type PostSpec struct {
	PostAt time.Time `json:"post_at"`
	Text   string    `json:"text"`
}

// PostStatus holds the status of the resource. E.g. if an
// operator processed the resource and what the result was.
type PostStatus struct {
	PostedOnTwitter   bool `json:"posted_on_twitter"`
	PostedOnFacebook  bool `json:"posted_on_facebook"`
	PostedOnInstagram bool `json:"posted_on_instagram"`
}
