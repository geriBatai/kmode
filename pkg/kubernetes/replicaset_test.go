package kubernetes_test

import (
	"reflect"
	"testing"

	"github.com/geriBatai/kmode/pkg/kubernetes"
)

func TestReplicaSetInit(t *testing.T) {
	data, err := buildLuaObject(`kubernetes.ReplicaSet()`)
	if err != nil {
		t.Errorf("error running lua code: %v", err)
	}

	rs := data.Value.(*kubernetes.ReplicaSet)
	if rs.Kind != "ReplicaSet" {
		t.Errorf("Unexpected kind. Expected: ReplicaSet, Got: %v", rs.Kind)
	}
}

func TestReplicaSetClone(t *testing.T) {
	o1 := &kubernetes.ReplicaSet{}
	o2 := o1.Clone()
	if !reflect.DeepEqual(o1, o2) {
		t.Errorf("Clone failed")
	}
}
