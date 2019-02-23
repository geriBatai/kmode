package kubernetes_test

import (
	"reflect"
	"testing"

	"github.com/geriBatai/kmode/pkg/kubernetes"
)

func TestPodInit(t *testing.T) {
	data, err := buildLuaObject(`kubernetes.Pod()`)
	if err != nil {
		t.Errorf("error running lua code: %v", err)
	}

	pod := data.Value.(*kubernetes.Pod)
	if pod.Kind != "Pod" {
		t.Errorf("Unexpected kind. Expected: Pod, Got: %v", pod.Kind)
	}
}

func TestPodClone(t *testing.T) {
	o1 := &kubernetes.Pod{}
	o2 := o1.Clone()
	if !reflect.DeepEqual(o1, o2) {
		t.Errorf("Clone failed")
	}
}
