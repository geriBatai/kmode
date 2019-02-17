package kubernetes

import (
	"bytes"
	"encoding/gob"

	luar "github.com/geriBatai/gopher-luar"
	lua "github.com/yuin/gopher-lua"
	appsv1 "k8s.io/api/apps/v1"

	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"fmt"

	"k8s.io/kubernetes/pkg/kubectl/generate/versioned"
)

type Deployment struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	*appsv1.Deployment
}

func (s *Deployment) Copy() *Deployment {
	newobj := &Deployment{}
	buff := new(bytes.Buffer)
	enc := gob.NewEncoder(buff)
	dec := gob.NewDecoder(buff)
	enc.Encode(s)
	dec.Decode(newobj)

	return newobj
}

func newDeployment(L *lua.LState) int {
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
	obj := &Deployment{
		Kind:       "Deployment",
		APIVersion: "apps/v1",
		Deployment: o.(*appsv1.Deployment),
	}
	L.Push(luar.New(L, obj))
	return 1
}
