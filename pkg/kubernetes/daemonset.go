package kubernetes

import (
	appsv1 "k8s.io/api/apps/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DaemonSet is a wrapper around Kubernetes runtime.Object
// for lua
type DaemonSet struct {
	*appsv1.DaemonSet
}

// Clone returns a duplicate object. Used in lua as object::clone()
func (d *DaemonSet) Clone() Resource {
	return copyResource(d, &DaemonSet{})
}

func defaultDaemonSet(options map[string]interface{}) Resource {
	return &DaemonSet{
		DaemonSet: &appsv1.DaemonSet{
			TypeMeta: metav1.TypeMeta{
				Kind:       "DaemonSet",
				APIVersion: "apps/v1",
			},
		},
	}
}
