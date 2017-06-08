package command

import (
	"github.com/spf13/cobra"
)

// Constant and data type/structure definitions
const version = "0.1.0"


//The version command prints this service.
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version.",
	Long:  "The version of the npa ",
	Run: func(cmd *cobra.Command, args []string) {
		println("npad version ", version)
	},
}
