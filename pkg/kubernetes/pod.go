package kubernetes

import (
	v1 "k8s.io/api/core/v1"
)

type Pod struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	*v1.Pod
}

func (p *Pod) Clone() KubernetesResource {
	return copyResource(p, &Pod{})
}

func defaultPod() KubernetesResource {
	return &Pod{
		Kind:       "Pod",
		APIVersion: "v1",
		Pod:        &v1.Pod{},
	}
}
