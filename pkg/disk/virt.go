package disk

import (
	"fmt"

	"github.com/zulcss/stx-installer/pkg/utils"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func CreateVirtDisk() {
	log.Info("Creating virtual disk")

	disk := viper.GetString("stx.instance.disk.device")
	size := viper.GetString("stx.instance.disk.size")

	// Create a file using truncate
	_, err := utils.SH(fmt.Sprintf("truncate --size %s %s", size, disk))
	if err != nil {
		fmt.Println("Truncate failed - %s", err)
	}

	// Create a GPT label for the disk, GPT is used
	// otherwise systemd-repart will fail.
	CreatePartition(disk)

}
