package images

import (
	"fmt"
//	"os"
//	"context"

	"github.com/zulcss/stx-installer/pkg/constants"

	 log "github.com/sirupsen/logrus"
//	 getter "github.com/hashicorp/go-getter"
	 "github.com/spf13/viper"
)


func FetchImage() {
	log.Info("Fetching image from download server")
	var url, version, img string


	url = viper.GetString("stx.image.url")
	if url == "" {
		log.Error("Unable to determine download server")
		return
	}
	version = viper.GetString("stx.version")
	if version == "" {
		log.Error("Unable to determine version to download")
		return
	}

	img = fmt.Sprintf("edgeos-%s.tar.gz", version)
	fmt.Println(constants.StarlingxVersion()["current"])
	fmt.Println(img)
	//download = "http://" + url + "/" + "/" 

/*
	client := &getter.Client{
		Src: "http://192.168.0.16/stx/",
	}
	if err := client.Get(); err == nil {
		fmt.Errorf("failed to download syncthing from %s: %s", client.Src, err)
	}

*/
}

func ImagePreFlight(version) {
	fmt.Println(version)
}
