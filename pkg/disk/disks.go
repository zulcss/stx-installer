package disk

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/zulcss/stx-installer/pkg/utils"
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

func CreatePartition(disk string) {
	out, err := utils.SH(fmt.Sprintf("parted %s gpt", disk))
	if err != nil {
		log.Error("Parted failed %s - %s", err, out)
	}
}
