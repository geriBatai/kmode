package luavm

import (
	"os"

	"github.com/geriBatai/kmode/pkg/kubernetes"
	lua "github.com/yuin/gopher-lua"
)

type VM struct {
	*lua.LState
}

// New returns a new VM object with kubernetes library loaded
func New() *VM {
	return &VM{
		LState: lua.NewState(),
	}
}

// Run loads variables and runs code in main filename
func (vm *VM) Run(vars, filename string) error {
	if err := vm.run(vars); err != nil {
		return err
	}
	vm.PreloadModule("kubernetes", kubernetes.Loader)
	return vm.run(filename)
}

func (vm *VM) run(filename string) error {
	if filename != "" {
		if _, err := os.Stat(filename); err != nil {
			return err
		}

		// Must be a better way
		// os.Setenv("LUA_PATH", "/.../?.lua")
		return vm.DoFile(filename)
	}
	return nil
}

// KubernetesGlobals loads all global kubernetes resources
// from the lua state and returns them as a map
func (vm *VM) KubernetesGlobals() map[string]kubernetes.Resource {
	table := vm.GetGlobal("_G").(*lua.LTable)

	var retval = map[string]kubernetes.Resource{}

	table.ForEach(func(a, b lua.LValue) {
		if b.Type() == lua.LTUserData {
			ud := b.(*lua.LUserData)

			switch ud.Value.(type) {
			case kubernetes.Resource:
				key := string(a.(lua.LString))
				retval[key] = ud.Value.(kubernetes.Resource)
			}
		}
	})
	return retval
}
