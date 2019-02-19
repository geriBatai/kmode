package kubernetes

import (
	v1 "k8s.io/api/core/v1"
)

type ServiceAccount struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	*v1.ServiceAccount
}

func (s *ServiceAccount) Clone() KubernetesResource {
	return copyResource(s, &ServiceAccount{})
}

func defaultServiceAccount() KubernetesResource {
	o := &v1.ServiceAccount{}
	return &ServiceAccount{
		Kind:           "ServiceAccount",
		APIVersion:     "v1",
		ServiceAccount: o,
	}
}
