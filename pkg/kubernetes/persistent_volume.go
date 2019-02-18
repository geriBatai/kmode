package kubernetes

import (
	v1 "k8s.io/api/core/v1"
)

type PersistentVolume struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	*v1.PersistentVolume
}

func (p *PersistentVolume) Copy() KubernetesResource {
	return cloneResource(p, &PersistentVolume{})
}

func defaultPersistentVolume() KubernetesResource {
	return &PersistentVolume{
		Kind:             "PersistentVolume",
		APIVersion:       "v1",
		PersistentVolume: &v1.PersistentVolume{},
	}
}
