package kubernetes

import (
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// StatefulSet is a wrapper around Kubernetes runtime.Object
// for lua
type StatefulSet struct {
	*appsv1.StatefulSet
}

// Clone returns a duplicate object. Used in lua as object::clone()
func (s *StatefulSet) Clone() Resource {
	return copyResource(s, &StatefulSet{})
}

func defaultStatefulSet(options map[string]interface{}) Resource {

	return &StatefulSet{
		StatefulSet: &appsv1.StatefulSet{
			TypeMeta: metav1.TypeMeta{
				Kind:       "StatefulSet",
				APIVersion: "apps/v1",
			},
		},
	}
}
