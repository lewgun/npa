package command

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"fmt"
)

// The root command describes the service and defaults to printing the
// help message.
var Root = &cobra.Command{
	Use:   "npad",
	Short: "a spider",
	Long:  `a spider which is named "npa"`,
	Run: func(cmd *cobra.Command, args []string) {
		entry(cmd)
	},
}

func init() {
	Root.AddCommand(versionCmd)
	Root.Flags().String("conf", "configs/npad.toml", "the path to the config file")

}


func entry(cmd *cobra.Command) {
	conf := cmd.Flags().Lookup("conf")
	viper.SetConfigType("toml")
	viper.SetConfigFile(conf.Value.String())
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
	fmt.Println("hello npa", viper.GetString("core.run.mode"))
}