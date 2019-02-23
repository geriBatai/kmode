package kubernetes_test

import (
	"reflect"
	"testing"

	"github.com/geriBatai/kmode/pkg/kubernetes"
)

func TestServiceInit(t *testing.T) {
	data, err := buildLuaObject(`kubernetes.Service()`)
	if err != nil {
		t.Errorf("error running lua code: %v", err)
	}

	svc := data.Value.(*kubernetes.Service)
	if svc.Kind != "Service" {
		t.Errorf("Unexpected kind. Expected: Service, Got: %v", svc.Kind)
	}
}

func TestServiceClone(t *testing.T) {
	o1 := &kubernetes.Service{}
	o2 := o1.Clone()
	if !reflect.DeepEqual(o1, o2) {
		t.Errorf("Clone failed")
	}
}
