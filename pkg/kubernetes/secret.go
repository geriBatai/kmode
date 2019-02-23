package kubernetes

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/kubernetes/pkg/kubectl/generate/versioned"
)

// Secret is a wrapper around Kubernetes runtime.Object
// for lua
type Secret struct {
	*v1.Secret
}

// Clone returns a duplicate object. Used in lua as object::clone()
func (s *Secret) Clone() Resource {
	return copyResource(s, &Secret{})
}

func defaultSecret(options map[string]interface{}) Resource {
	generator := versioned.SecretGeneratorV1{}
	options["name"] = "secret"
	o, err := generator.Generate(options)
	if err != nil {
		fmt.Printf("ERROR generating Secret resource: %v\n", err)
	}

	gvk := schema.GroupVersionKind{
		Kind:    "Secret",
		Version: "v1",
	}
	o.GetObjectKind().SetGroupVersionKind(gvk)

	return &Secret{
		Secret: o.(*v1.Secret),
	}
}
