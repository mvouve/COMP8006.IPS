package main

import (
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/Unknwon/goconfig"
)

const configIni string = "config.ini"

var configure *goconfig.ConfigFile

type manifestType struct {
	filePos int64
	bans    map[time.Time]string
}

func man() {
	var manifest manifestType

	configure, err := goconfig.LoadConfigFile(configIni)
	if os.IsNotExist(err) {
		log.Fatal("config.ini should be in the same directory as this is run from")
	}

	filePath, _ := configure.GetValue("manifest", "file")
	if manifestFile, err := os.Open(filePath); os.IsNotExist(err) {
		manifest = initManifest()
	} else {
		manifest = loadManifest(manifestFile)
	}
	checkBans(&manifest.bans)
	checkSecure(&manifest.filePos)

}

func initManifest() manifestType {
	return manifestType{}
}

func loadManifest(file *os.File) manifestType {
	return manifestType{}
}

func checkBans(currentBans *map[time.Time]string) {
	for expiry, ip := range *currentBans {
		if time.Now().Before(expiry) {
			dropBan(ip)
			delete(*currentBans, expiry)
		}
	}
}

func dropBan(ip string) {
	for _, chain := range []string{"INPUT", "OUTPUT", "FORWARD"} {
		exec.Command("iptables", "-D", chain, "-s", ip, "-j", "DROP")
	}
}

func addBan(ip string) {
	for _, chain := range []string{"INPUT", "OUTPUT", "FORWARD"} {
		exec.Command("iptables", "-A", chain, "-s", ip, "-j", "DROP")
	}
}

func checkSecure(filePos *int64) {
	filepath, _ := configure.GetValue("auth", "file")
	logFile, err := os.Open(filepath)
	if os.IsNotExist(err) {
		log.Fatalln("secure file not found, if rsyslog is installed check the ini")
	}
	logFile.Seek(*filePos, 0)
}
