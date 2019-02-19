package kubernetes_test

import (
	"reflect"
	"testing"

	"github.com/geriBatai/kmode/pkg/kubernetes"
)

func TestSecretInit(t *testing.T) {
	data, err := buildLuaObject(`kubernetes.Secret()`)
	if err != nil {
		t.Errorf("error running lua code: %v", err)
	}

	secret := data.Value.(*kubernetes.Secret)
	if secret.Kind != "Secret" {
		t.Errorf("Unexpected kind. Expected: Secret, Got: %v", secret.Kind)
	}

}

func TestSecretClone(t *testing.T) {
	s1 := &kubernetes.Secret{}
	s2 := s1.Clone()
	if !reflect.DeepEqual(s1, s2) {
		t.Errorf("Clone failed")
	}
}
