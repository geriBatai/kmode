package kubernetes

import (
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ReplicaSet is a wrapper around Kubernetes runtime.Object
// for lua
type ReplicaSet struct {
	*appsv1.ReplicaSet
}

// Clone returns a duplicate object. Used in lua as object::clone()
func (r *ReplicaSet) Clone() Resource {
	return copyResource(r, &ReplicaSet{})
}

func defaultReplicaSet(options map[string]interface{}) Resource {
	o := &appsv1.ReplicaSet{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ReplicaSet",
			APIVersion: "apps/v1",
		},
	}
	return &ReplicaSet{
		ReplicaSet: o,
	}
}
