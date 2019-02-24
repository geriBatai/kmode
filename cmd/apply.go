package cmd

import "github.com/spf13/cobra"

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply generated objects to the Kubernetes cluster",
	Run:   runApply,
}

var kubeconfig string

func runApply(cmd *cobra.Command, args []string) {
	// TBC
}
