package kubernetes

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ServiceAccount is a wrapper around Kubernetes runtime.Object
// for lua
type ServiceAccount struct {
	*v1.ServiceAccount
}

// Clone returns a duplicate object. Used in lua as object::clone()
func (s *ServiceAccount) Clone() Resource {
	return copyResource(s, &ServiceAccount{})
}

func defaultServiceAccount() Resource {
	o := &v1.ServiceAccount{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ServiceAccount",
			APIVersion: "v1",
		},
	}
	return &ServiceAccount{
		ServiceAccount: o,
	}
}
