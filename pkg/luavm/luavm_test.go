package luavm_test

import (
	"testing"

	"github.com/geriBatai/kmode/pkg/luavm"
)

func TestRun(t *testing.T) {
	runs := []struct {
		variables string
		code      string
		success   bool
	}{{
		``,
		`local kubernetes = require("kubernetes"); d = kubernetes.Deployment()`,
		true,
	},
		{
			`a = 59`,
			`d = kubernetes.Deployment()`,
			false,
		},
	}

	vm := luavm.New(nil)

	for _, run := range runs {
		e := vm.Run(run.variables, run.code)
		if (e == nil) != run.success {
			t.Errorf("expected e to be %t, got %v", run.success, e)
		}
	}

}

func TestKubernetesGlobals(t *testing.T) {
	code := `local kubernetes = require("kubernetes"); d = kubernetes.Deployment{name="test"}`

	vm := luavm.New(nil)
	vm.Run("", code)
	globals := vm.KubernetesGlobals()

	if len(globals) != 1 {
		t.Errorf("Expected 1 global, got: %v", len(globals))
	}
}
