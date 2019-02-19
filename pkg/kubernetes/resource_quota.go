package kubernetes

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ResourceQuota is a wrapper around Kubernetes runtime.Object
// for lua
type ResourceQuota struct {
	*v1.ResourceQuota
}

// Clone returns a duplicate object. Used in lua as object::clone()
func (r *ResourceQuota) Clone() Resource {
	return copyResource(r, &ResourceQuota{})
}

func defaultResourceQuota() Resource {
	o := &v1.ResourceQuota{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ResourceQuota",
			APIVersion: "v1",
		},
	}
	return &ResourceQuota{
		ResourceQuota: o,
	}
}
