package cmd

import (
	"fmt"

	"github.com/jmau111-org/good/libgood"
	"github.com/spf13/cobra"
)

var versionCmd *cobra.Command

func runCmdVersion(cmd *cobra.Command, args []string) error {
	fmt.Printf("%s\n\n", libgood.VERSION)
	return nil
}

func init() {
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Shows the current version",
		RunE:  runCmdVersion,
	}

	checkCmd.AddCommand(versionCmd)
}
