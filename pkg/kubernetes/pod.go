package kubernetes

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Pod is a wrapper around Kubernetes runtime.Object
// for lua
type Pod struct {
	*v1.Pod
}

// Clone returns a duplicate object. Used in lua as object::clone()
func (p *Pod) Clone() Resource {
	return copyResource(p, &Pod{})
}

func defaultPod(options map[string]interface{}) Resource {
	return &Pod{
		Pod: &v1.Pod{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Pod",
				APIVersion: "v1",
			},
		},
	}
}
