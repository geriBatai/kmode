package kubernetes

import (
	v1 "k8s.io/api/core/v1"
)

type ReplicationController struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	*v1.ReplicationController
}

func (r *ReplicationController) Copy() KubernetesResource {
	return cloneResource(r, &ReplicationController{})
}

func defaultReplicationController() KubernetesResource {
	o := &v1.ReplicationController{}
	return &ReplicationController{
		Kind:                  "ReplicationController",
		APIVersion:            "v1",
		ReplicationController: o,
	}
}
