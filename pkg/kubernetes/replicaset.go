package kubernetes

import (
	appsv1 "k8s.io/api/apps/v1"
)

type ReplicaSet struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	*appsv1.ReplicaSet
}

func (r *ReplicaSet) Clone() KubernetesResource {
	return copyResource(r, &ReplicaSet{})
}

func defaultReplicaSet() KubernetesResource {
	o := &appsv1.ReplicaSet{}
	return &ReplicaSet{
		Kind:       "ReplicaSet",
		APIVersion: "apps/v1",
		ReplicaSet: o,
	}
}
