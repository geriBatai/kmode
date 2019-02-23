package kubernetes_test

import (
	"reflect"
	"testing"

	"github.com/geriBatai/kmode/pkg/kubernetes"
)

func TestDaemonSetInit(t *testing.T) {
	data, err := buildLuaObject(`kubernetes.DaemonSet()`)
	if err != nil {
		t.Errorf("error running lua code: %v", err)
	}

	daemonset := data.Value.(*kubernetes.DaemonSet)
	if daemonset.Kind != "DaemonSet" {
		t.Errorf("Unexpected kind. Expected: DaemonSet, Got: %v", daemonset.Kind)
	}
}

func TestDaemonSetClone(t *testing.T) {
	o1 := &kubernetes.DaemonSet{}
	o2 := o1.Clone()
	if !reflect.DeepEqual(o1, o2) {
		t.Errorf("Clone failed")
	}
}
