package cmd

import (
	"fmt"
	"os"

	"github.com/zulcss/stx-installer/internal"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "stx-installer",
	Short: "StarlingX Instaler",
	// TODO: write something here
	Long: "",
	Run: func(cmd *cobra.Command, args []string) { 
		internal.ReadConfig()

		fmt.Printf("%s\n", viper.GetString("image.url"))
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&internal.ConfigFile, "config", "C", "/etc/stx/installer.yaml", "configuration file")
}


