package kubernetes

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
	"k8s.io/kubernetes/pkg/kubectl/generate/versioned"
)

type Secret struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	*v1.Secret
}

func (s *Secret) Copy() KubernetesResource {
	return cloneResource(s, &Secret{})
}

func defaultSecret() KubernetesResource {
	generator := versioned.SecretGeneratorV1{}
	opts := map[string]interface{}{}
	opts["name"] = "secret"
	o, err := generator.Generate(opts)
	if err != nil {
		fmt.Printf("ERROR generating Secret resource: %v\n", err)
	}
	return &Secret{
		Kind:       "Secret",
		APIVersion: "v1",
		Secret:     o.(*v1.Secret),
	}
}
