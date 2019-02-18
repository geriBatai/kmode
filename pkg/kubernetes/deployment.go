package kubernetes

import (
	"fmt"

	appsv1 "k8s.io/api/apps/v1"

	"k8s.io/kubernetes/pkg/kubectl/generate/versioned"
)

type Deployment struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	*appsv1.Deployment
}

func (d *Deployment) Copy() KubernetesResource {
	return cloneResource(d, &Deployment{})
}

func defaultDeployment() KubernetesResource {
	generator := versioned.DeploymentBasicAppsGeneratorV1{
		BaseDeploymentGenerator: versioned.BaseDeploymentGenerator{
			Name:   "default",
			Images: []string{"nginx:default"},
		},
	}

	o, err := generator.StructuredGenerate()
	if err != nil {
		fmt.Printf("ERROR generating Deployment resource: %v\n", err)
	}
	return &Deployment{
		Kind:       "Deployment",
		APIVersion: "apps/v1",
		Deployment: o.(*appsv1.Deployment),
	}
}
