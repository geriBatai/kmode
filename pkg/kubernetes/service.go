package kubernetes

import (
	"fmt"

	v1 "k8s.io/api/core/v1"

	//  "github.com/kubernetes/kubernetes/pkg/kubectl/generate/versioned"
	"k8s.io/kubernetes/pkg/kubectl/generate/versioned"
)

type Service struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	*v1.Service
}

func (s *Service) Copy() KubernetesResource {
	return cloneResource(s, &Service{})
}

func defaultService() KubernetesResource {
	generator := versioned.ServiceGeneratorV1{}
	opts := map[string]interface{}{}
	opts["default-name"] = "svc"
	opts["selector"] = "name=svc"
	opts["port"] = "80"

	s, err := generator.Generate(opts)
	if err != nil {
		fmt.Printf("ERROR generating Service resource: %v\n", err)
	}
	return &Service{
		Kind:       "Service",
		APIVersion: "v1",
		Service:    s.(*v1.Service),
	}
}
