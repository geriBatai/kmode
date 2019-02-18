package kubernetes_test

import (
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
