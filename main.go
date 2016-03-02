package main

import (
	"log"
	"os"

	"github.com/Unknwon/goconfig"
)

var configure *goconfig.ConfigFile

func main() {
	var ips manifestType
	var err error

	configure, err = goconfig.LoadConfigFile(configIni)
	if os.IsNotExist(err) {
		log.Fatal("config.ini should be in the same directory as this is run from")
	} else if err != nil {
		log.Fatal("error opening config: ", err)
	}

	filePath := configure.MustValue("manifest", "file", "/root/go/src/github.com/mvouve/COMP8006.IPS/manifest")
	if manifestFile, err := os.Open(filePath); os.IsNotExist(err) {
		ips = initManifest()
	} else {
		ips = loadManifest(manifestFile)
		manifestFile.Close()
	}

	ips.checkBans()
	ips.checkSecure()
	ips.checkEvents()
	ips.save(filePath)
}
