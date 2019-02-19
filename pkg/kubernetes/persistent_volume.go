package kubernetes

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PersistentVolume is a wrapper around Kubernetes runtime.Object
// for lua
type PersistentVolume struct {
	*v1.PersistentVolume
}

// Clone returns a duplicate object. Used in lua as object::clone()
func (p *PersistentVolume) Clone() Resource {
	return copyResource(p, &PersistentVolume{})
}

func defaultPersistentVolume() Resource {
	return &PersistentVolume{
		PersistentVolume: &v1.PersistentVolume{
			TypeMeta: metav1.TypeMeta{
				Kind:       "PersistentVolume",
				APIVersion: "v1",
			},
		},
	}
}
