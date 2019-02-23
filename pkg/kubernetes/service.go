package kubernetes

import (
	"fmt"

	v1 "k8s.io/api/core/v1"

	//  "github.com/kubernetes/kubernetes/pkg/kubectl/generate/versioned"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/kubernetes/pkg/kubectl/generate/versioned"
)

// Service is a wrapper around Kubernetes runtime.Object
// for lua
type Service struct {
	*v1.Service
}

// Clone returns a duplicate object. Used in lua as object::clone()
func (s *Service) Clone() Resource {
	return copyResource(s, &Service{})
}

func defaultService(options map[string]interface{}) Resource {
	generator := versioned.ServiceGeneratorV1{}
	opts := map[string]interface{}{}
	opts["default-name"] = "svc"
	opts["selector"] = "name=svc"
	opts["port"] = "80"

	s, err := generator.Generate(opts)
	if err != nil {
		fmt.Printf("ERROR generating Service resource: %v\n", err)
	}

	gvk := schema.GroupVersionKind{
		Kind:    "Service",
		Version: "v1",
	}
	s.GetObjectKind().SetGroupVersionKind(gvk)
	return &Service{
		Service: s.(*v1.Service),
	}
}
