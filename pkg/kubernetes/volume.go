package kubernetes

import (
	v1 "k8s.io/api/core/v1"
)

type Volume struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	*v1.Volume
}

func (v *Volume) Copy() KubernetesResource {
	return cloneResource(v, &Volume{})
}

func defaultVolume() KubernetesResource {
	return &Volume{
		Kind:       "Secret",
		APIVersion: "v1",
		Volume:     &v1.Volume{},
	}
}
