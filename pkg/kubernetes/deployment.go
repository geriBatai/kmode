package kubernetes

import (
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/kubernetes/pkg/kubectl/generate/versioned"
)

// Deployment is a wrapper around Kubernetes runtime.Object
// for lua
type Deployment struct {
	*appsv1.Deployment
}

// Clone returns a duplicate object. Used in lua as object::clone()
func (d *Deployment) Clone() Resource {
	return copyResource(d, &Deployment{})
}

func defaultDeployment(options map[string]interface{}) Resource {
	deploymentName := "default"
	if options["name"] != nil {
		deploymentName = options["name"].(string)
	}
	// deploymentNamespace := options["namespace"].(string)

	generator := versioned.DeploymentBasicAppsGeneratorV1{
		BaseDeploymentGenerator: versioned.BaseDeploymentGenerator{
			Name:   deploymentName,
			Images: []string{"nginx:default"},
		},
	}

	o, err := generator.StructuredGenerate()
	if err != nil {
		fmt.Printf("ERROR generating Deployment resource: %v\n", err)
	}
	gvk := schema.GroupVersionKind{
		Kind:    "Deployment",
		Version: "apps/v1",
	}
	o.GetObjectKind().SetGroupVersionKind(gvk)
	return &Deployment{
		Deployment: o.(*appsv1.Deployment),
	}
}
