package cmd

import (
	//"fmt"
	"os"

	"github.com/zulcss/stx-installer/internal"
	"github.com/zulcss/stx-installer/pkg/installer"

	"github.com/spf13/cobra"
	//"github.com/spf13/viper"
	log "github.com/sirupsen/logrus"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "stx-installer",
	Short: "StarlingX Instaler",
	// TODO: write something here
	Long: "",
	Run: func(cmd *cobra.Command, args []string) { 
		installer.Install()
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

	log.SetLevel(log.InfoLevel)
}


