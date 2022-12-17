package installer

import (
	//"fmt"
	"github.com/zulcss/stx-installer/internal"
	"github.com/zulcss/stx-installer/pkg/images"

	log "github.com/sirupsen/logrus"
//	"github.com/spf13/viper"
)

// Main installer
func Install() {
	log.Info("Installing StarlingX")

	// Read the configuration file
	internal.ReadConfig()

	// Fetch the image from the download server
	images.FetchImage()
	//fmt.Println(viper.GetString("image.url"))


}
