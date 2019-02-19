package kubernetes

import (
	appsv1 "k8s.io/api/apps/v1"
)

type StatefulSet struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	*appsv1.StatefulSet
}

func (s *StatefulSet) Clone() KubernetesResource {
	return copyResource(s, &StatefulSet{})
}

func defaultStatefulSet() KubernetesResource {

	return &StatefulSet{
		Kind:        "StatefulSet",
		APIVersion:  "apps/v1",
		StatefulSet: &appsv1.StatefulSet{},
	}
}
