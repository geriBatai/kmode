package kubernetes_test

import (
	"reflect"
	"testing"

	"github.com/geriBatai/kmode/pkg/kubernetes"
)

func TestStatefulSetInit(t *testing.T) {
	data, err := buildLuaObject(`kubernetes.StatefulSet()`)
	if err != nil {
		t.Errorf("error running lua code: %v", err)
	}

	set := data.Value.(*kubernetes.StatefulSet)
	if set.Kind != "StatefulSet" {
		t.Errorf("Unexpected kind. Expected: StatefulSet, Got: %v", set.Kind)
	}
}

func TestStatefulSetClone(t *testing.T) {
	o1 := &kubernetes.StatefulSet{}
	o2 := o1.Clone()
	if !reflect.DeepEqual(o1, o2) {
		t.Errorf("Clone failed")
	}
}
