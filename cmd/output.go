package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/geriBatai/kmode/pkg/luavm"
	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"
)

var outputCmd = &cobra.Command{
	Use:   "output",
	Short: "Print generated objects to the standard output",
	Run:   runOutput,
}

func runOutput(cmd *cobra.Command, args []string) {
	vars := readFile(varFile)
	contents := readFile(filename)
	modulePath := filepath.Dir(filename) + "/?.lua"

	vm := luavm.New(&luavm.Options{
		LuaPath: modulePath,
	})
	defer vm.Close()

	if err := vm.Run(vars, contents); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}

	for _, v := range vm.KubernetesGlobals() {
		o, err := yaml.Marshal(v)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("---\n%s\n", o)
	}
}

func readFile(filename string) string {
	if filename != "" {
		if _, err := os.Stat(filename); err != nil {
			fmt.Fprintf(os.Stderr, "ERROR %v\n", err)
			os.Exit(1)
		}

		c, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
			os.Exit(1)
		}
		return string(c)
	}
	return ""
}
