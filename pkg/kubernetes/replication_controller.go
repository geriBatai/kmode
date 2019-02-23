package kubernetes

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ReplicationController is a wrapper around Kubernetes runtime.Object
// for lua
type ReplicationController struct {
	*v1.ReplicationController
}

// Clone returns a duplicate object. Used in lua as object::clone()
func (r *ReplicationController) Clone() Resource {
	return copyResource(r, &ReplicationController{})
}

func defaultReplicationController(options map[string]interface{}) Resource {
	o := &v1.ReplicationController{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ReplicationController",
			APIVersion: "v1",
		},
	}
	return &ReplicationController{
		ReplicationController: o,
	}
}
