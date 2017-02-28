package cli

import (
	"fmt"

	"github.com/tonglil/versioning"

	"github.com/spf13/cobra"
)

// CobraVersionCmd represents the version command
// Add this to incorporate this command into your Cobra CLI application:
// import "github.com/tonglil/versioning/cli"
// func init() {
// 	RootCmd.AddCommand(cli.CobraVersionCmd)
// }
var CobraVersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("version %s\n", versioning.Get())
	},
}
