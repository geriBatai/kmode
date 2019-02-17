package cmd

import (
	"fmt"
	"os"

	"github.com/geriBatai/kmode/pkg/kubernetes"
	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"
	lua "github.com/yuin/gopher-lua"
)

var rootCmd = &cobra.Command{
	Use: "kmode",
	Run: runRoot,
}

var filename string
var varFile string

func init() {
	rootCmd.PersistentFlags().StringVar(&filename, "filename", "filename.lua", "kmode filename")
	rootCmd.PersistentFlags().StringVar(&varFile, "var-file", "", "variable file")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runRoot(cmd *cobra.Command, args []string) {
	L := lua.NewState()
	defer L.Close()

	runLuaFile(L, varFile)
	L.PreloadModule("kubernetes", kubernetes.Loader)
	runLuaFile(L, filename)

	table := L.GetGlobal("_G").(*lua.LTable)
	table.ForEach(func(a, b lua.LValue) {
		if b.Type() == lua.LTUserData {
			marshalValue(b)
		}
	})
}

func runLuaFile(L *lua.LState, filename string) error {
	if filename != "" {
		if _, err := os.Stat(filename); err != nil {
			return fmt.Errorf("Error reading file %s: %s", filename, err)
		}
		return L.DoFile(filename)
	}
	return nil
}

func marshalValue(val lua.LValue) {
	ud := val.(*lua.LUserData)

	res, err := yaml.Marshal(ud.Value)
	if err != nil {
		fmt.Printf("Error marshaling to yaml: %s\n", err.Error())
	}
	fmt.Printf("%v\n----\n", string(res))
}
