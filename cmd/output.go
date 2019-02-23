package cmd

import (
	"fmt"

	"github.com/geriBatai/kmode/pkg/kubernetes"
	"github.com/spf13/cobra"
	lua "github.com/yuin/gopher-lua"
)

var outputCmd = &cobra.Command{
	Use:   "output",
	Short: "Print generated objects to the standard output",
	Run:   runOutput,
}

func runOutput(cmd *cobra.Command, args []string) {
	// Must be a better way
	// os.Setenv("LUA_PATH", "/.../?.lua")
	L := lua.NewState()
	defer L.Close()

	if err := runLuaFile(L, varFile); err != nil {
		fmt.Printf("ERROR: %v", err)
	}
	L.PreloadModule("kubernetes", kubernetes.Loader)
	if err := runLuaFile(L, filename); err != nil {
		fmt.Printf("ERROR: %v\n", err)
		//errors.Wrap(err, 1).ErrorStack())
	}

	table := L.GetGlobal("_G").(*lua.LTable)
	table.ForEach(func(a, b lua.LValue) {
		if b.Type() == lua.LTUserData {
			marshalValue(b)
		}
	})
}
