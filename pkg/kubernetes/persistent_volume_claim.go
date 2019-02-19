package kubernetes

import (
	v1 "k8s.io/api/core/v1"
)

type PersistentVolumeClaim struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	*v1.PersistentVolumeClaim
}

func (p *PersistentVolumeClaim) Clone() KubernetesResource {
	return copyResource(p, &PersistentVolumeClaim{})
}

func defaultPersistentVolumeClaim() KubernetesResource {
	return &PersistentVolumeClaim{
		Kind:                  "PersistentVolumeClaim",
		APIVersion:            "v1",
		PersistentVolumeClaim: &v1.PersistentVolumeClaim{},
	}
}
