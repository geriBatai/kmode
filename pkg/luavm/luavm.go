package luavm

import (
	"os"

	"github.com/geriBatai/kmode/pkg/kubernetes"
	lua "github.com/yuin/gopher-lua"
)

// VM is a wrapper around lua.LState
type VM struct {
	*lua.LState
}

type Options struct {
	LuaPath string
}

// New returns a new VM object with kubernetes library loaded
func New(options Options) *VM {
	os.Setenv("LUA_PATH", options.LuaPath)
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
