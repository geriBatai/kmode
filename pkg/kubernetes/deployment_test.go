package kubernetes_test

import (
	"reflect"
	"testing"

	"github.com/geriBatai/kmode/pkg/kubernetes"
)

func TestDeploymentInit(t *testing.T) {
	data, err := buildLuaObject(`kubernetes.Deployment()`)
	if err != nil {
		t.Errorf("error running lua code: %v", err)
	}

	deployment := data.Value.(*kubernetes.Deployment)
	if deployment.Kind != "Deployment" {
		t.Errorf("Unexpected kind. Expected: Deployment, Got: %v", deployment.Kind)
	}
}

func TestDeploymentClone(t *testing.T) {
	o1 := &kubernetes.Deployment{}
	o2 := o1.Clone()
	if !reflect.DeepEqual(o1, o2) {
		t.Errorf("Clone failed")
	}
}
