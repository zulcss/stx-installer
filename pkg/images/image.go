package images

import (
	"context"

	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/zulcss/stx-installer/pkg/constants"

	getter "github.com/hashicorp/go-getter"
	"github.com/spf13/viper"
)

func FetchImage() {
	log.Info("Fetching image from download server")

	version, url := ImagePreFlightCheck()
	img := fmt.Sprintf("http://%s/stx/edgeos-%s.tar.gz", url, version)
	DownloadImage(img)
}

func DownloadImage(image string) {
	// Download the image from a server
	log.Info("Downloading image")

	client := &getter.Client{
		Ctx:  context.Background(),
		Dst:  constants.OEMImageDir,
		Src:  image,
		Dir:  true,
		Mode: getter.ClientModeDir,
	}

	if err := client.Get(); err != nil {
		log.Error("Failed to download image: %s ", err)
	}
}

func ImagePreFlightCheck() (version string, url string) {
	// Get the starlingx and check for valid
	version = viper.GetString("stx.version")
	if version == "" && constants.StarlingxVersion()[version] == version {
		log.Error("Unable to deletermine image version")
	}

	// Check the image URL
	url = viper.GetString("stx.image.server")
	if url == "" {
		log.Error("Unable to determine URL.")
	}

	// TODO: verify md5sum of the tarball
	return
}
