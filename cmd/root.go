package cmd

import (
	"fmt"
	"os"
	"reflect"

	"github.com/geriBatai/kmode/pkg/kubernetes"
	"github.com/kubernetes/cli-runtime/pkg/genericclioptions/printers"
	"github.com/spf13/cobra"
	lua "github.com/yuin/gopher-lua"
	"k8s.io/apimachinery/pkg/runtime"
)

var rootCmd = &cobra.Command{
	Use: "kmode",
}

var filename string
var varFile string

func init() {
	rootCmd.PersistentFlags().StringVar(&filename, "filename", "filename.lua", "kmode filename")
	rootCmd.PersistentFlags().StringVar(&varFile, "var-file", "", "variable file")
	rootCmd.AddCommand(outputCmd)
	rootCmd.AddCommand(applyCmd)
}

// Execute is the entry point for a kmode
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func marshalValue(val lua.LValue) {
	ud := val.(*lua.LUserData)

	switch ud.Value.(type) {
	case kubernetes.Resource:
		printer := &printers.YAMLPrinter{}
		fmt.Printf("---\n")
		err := printer.PrintObj(ud.Value.(runtime.Object), os.Stdout)
		if err != nil {
			fmt.Printf("ERROR: %v", err)
		}
	default:
		fmt.Printf("TYPE: %v\n", reflect.TypeOf(ud.Value))
	}
}
