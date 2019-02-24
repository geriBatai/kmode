package cmd

import (
	"fmt"
	"os"

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
	vm := luavm.New()
	defer vm.Close()

	if err := vm.Run(varFile, filename); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v", err)
		os.Exit(1)
	}

	for _, v := range vm.KubernetesGlobals() {
		o, err := yaml.Marshal(v)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %v", err)
			os.Exit(1)
		}

		fmt.Printf("---\n%s\n", o)
	}
}
