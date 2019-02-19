package kubernetes

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PersistentVolumeClaim is a wrapper around Kubernetes runtime.Object
// for lua
type PersistentVolumeClaim struct {
	*v1.PersistentVolumeClaim
}

// Clone returns a duplicate object. Used in lua as object::clone()
func (p *PersistentVolumeClaim) Clone() Resource {
	return copyResource(p, &PersistentVolumeClaim{})
}

func defaultPersistentVolumeClaim() Resource {
	return &PersistentVolumeClaim{
		PersistentVolumeClaim: &v1.PersistentVolumeClaim{
			TypeMeta: metav1.TypeMeta{
				Kind:       "PersistentVolumeClaim",
				APIVersion: "v1",
			},
		},
	}
}
