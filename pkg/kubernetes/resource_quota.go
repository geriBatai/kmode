package kubernetes

import (
	v1 "k8s.io/api/core/v1"
)

type ResourceQuota struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	*v1.ResourceQuota
}

func (r *ResourceQuota) Clone() KubernetesResource {
	return copyResource(r, &ResourceQuota{})
}

func defaultResourceQuota() KubernetesResource {
	o := &v1.ResourceQuota{}
	return &ResourceQuota{
		Kind:          "ResourceQuota",
		APIVersion:    "v1",
		ResourceQuota: o,
	}
}
