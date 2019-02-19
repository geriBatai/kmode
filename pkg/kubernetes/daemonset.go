package kubernetes

import (
	appsv1 "k8s.io/api/apps/v1"
)

type DaemonSet struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	*appsv1.DaemonSet
}

func (d *DaemonSet) Clone() KubernetesResource {
	return copyResource(d, &DaemonSet{})
}

func defaultDaemonSet() KubernetesResource {
	return &DaemonSet{
		Kind:       "DaemonSet",
		APIVersion: "apps/v1",
		DaemonSet:  &appsv1.DaemonSet{},
	}
}
