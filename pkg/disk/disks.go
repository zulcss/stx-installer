package disk

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	//"github.com/zulcss/stx-installer/pkg/constants"
)

func CreateDisk() {
	log.Info("Creating disk for StarlingX install")

	disk := viper.GetString("stx.instance.disk.type")
	if disk == "" {
		log.Error("Unable to determine disk type")
	}
	switch disk {
	case "virt":
		CreateVirtDisk()
	case "metal":
		CreateMetalDisk()
	default:
		log.Error("Unable to determine disk type")

	}
}
