package kubernetes_test

import (
	"reflect"
	"testing"

	"github.com/geriBatai/kmode/pkg/kubernetes"
)

func TestReplicationControllerInit(t *testing.T) {
	data, err := buildLuaObject(`kubernetes.ReplicationController()`)
	if err != nil {
		t.Errorf("error running lua code: %v", err)
	}

	rc := data.Value.(*kubernetes.ReplicationController)
	if rc.Kind != "ReplicationController" {
		t.Errorf("Unexpected kind. Expected: ReplicationController, Got: %v", rc.Kind)
	}
}

func TestReplicationControllerClone(t *testing.T) {
	o1 := &kubernetes.ReplicationController{}
	o2 := o1.Clone()
	if !reflect.DeepEqual(o1, o2) {
		t.Errorf("Clone failed")
	}
}
