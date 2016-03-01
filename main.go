package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"time"

	"github.com/Unknwon/goconfig"
)

const configIni string = "/root/Music/COMP8006.IPS/config.ini"
const defaultTrys int = 5 // default ammount of tries allowed (overridden by config)
const defaultBan int = 24 // time in hours for default ban (overriden by config)

var configure *goconfig.ConfigFile

type manifestType struct {
	FilePos int64
	Bans    map[string]time.Time
	Events  map[string][]time.Time
}

func main() {
	var manifest manifestType
	var err error

	configure, err = goconfig.LoadConfigFile(configIni)
	if os.IsNotExist(err) {
		log.Fatal("config.ini should be in the same directory as this is run from")
	}

	filePath := configure.MustValue("manifest", "file", "manifest")
	if manifestFile, err := os.Open(filePath); os.IsNotExist(err) {
		manifest = initManifest()
	} else {
		manifest = loadManifest(manifestFile)
	}
	checkBans(&manifest.Bans)
	checkSecure(&manifest.FilePos, manifest.Events)
	checkEvents(manifest.Bans, manifest.Events)
	save(filePath, manifest)
}

func initManifest() manifestType {
	events := make(map[string][]time.Time)
	bans := make(map[string]time.Time)

	return manifestType{Events: events, Bans: bans}
}

func loadManifest(file *os.File) manifestType {
	stats, _ := file.Stat()
	buffer := make([]byte, stats.Size())
	file.Read(buffer)
	mani := manifestType{}
	err := json.Unmarshal(buffer, &mani)
	if err != nil {
		mani = initManifest()
		log.Println(err)
	}
	return mani
}

func checkBans(currentBans *map[string]time.Time) {
	for ip, expiry := range *currentBans {
		if time.Now().After(expiry) {
			dropBan(ip)
			delete(*currentBans, ip)
		}
	}
}

func checkSecure(filePos *int64, events map[string][]time.Time) {
	filepath := configure.MustValue("auth", "file", "/var/log/secure")
	logFile, err := os.Open(filepath)
	if os.IsNotExist(err) {
		log.Fatalln("secure file not found, if rsyslog is installed check the ini")
	} else if err != nil {
		log.Fatalln(err)
	}
	defer logFile.Close()

	if _, err = logFile.Seek(*filePos, os.SEEK_SET); err != nil {
		logFile.Seek(0, 0)
		*filePos = 0
	}
	logBuff := bufio.NewReader(logFile)
	for {
		line, err := logBuff.ReadString('\n')
		*filePos += int64(len(line))
		if err == io.EOF {
			break
		}
		if isFailedLogin(line) {
			if events[getIPfromString(line)] == nil {
				events[getIPfromString(line)] = make([]time.Time, 0, 1024)
			}
			events[getIPfromString(line)] = append(events[getIPfromString(line)][0:], getTimeFromString(line))
		}
	}

}

func checkEvents(bans map[string]time.Time, events map[string][]time.Time) {
	now, _ := time.Parse("Jan 02 15:04:05", time.Now().Format("Jan 02 15:04:05"))

	for ip := range events {
		recentEvents := make([]time.Time, 0, 128)
		for idx := 0; idx < len(events[ip]); idx++ {
			if now.Sub(events[ip][idx]) < (time.Minute * time.Duration(configure.MustInt("auth", "trace_time", 1))) { // more than a minute has passed.
				recentEvents = append(recentEvents[0:], events[ip][idx])
			}
		}
		if len(recentEvents) >= configure.MustInt("auth", "max_attempts", defaultTrys) {
			addBan(ip, bans)
		}
	}
}

func dropBan(ip string) {
	for _, chain := range []string{"INPUT", "OUTPUT", "FORWARD"} {
		exec.Command("iptables", "-D", chain, "-s", ip, "-j", "DROP").Run()
	}
}

func addBan(ip string, bans map[string]time.Time) {
	_, ok := bans[ip]
	if !ok {
		for _, chain := range []string{"INPUT", "OUTPUT", "FORWARD"} {
			exec.Command("iptables", "-A", chain, "-s", ip, "-j", "DROP").Run()
		}
	}
	bans[ip] = time.Now().Add(time.Hour * time.Duration(configure.MustInt("auth", "ban_time", defaultBan)))
}

func isFailedLogin(log string) bool {
	failed, _ := regexp.MatchString("Failed password", log)
	return failed
}

func getIPfromString(log string) string {
	regx := regexp.MustCompile(`[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}`)

	return regx.FindString(log)
}

func getTimeFromString(log string) time.Time {
	t, _ := time.Parse("Jan 02 15:04:05", getTimeStringFromString(log))

	return t
}

func getTimeStringFromString(line string) string {
	regx := regexp.MustCompile(`[A-Za-z]{3}\s\d{1,2}\s\d{2}:\d{2}:\d{2}`)

	return regx.FindString(line)
}

func save(f string, m manifestType) {
	data, err := json.Marshal(m)
	if err != nil {
		log.Fatalln(err)
	}
	file, _ := os.Create(f)
	defer file.Close()

	var out bytes.Buffer
	json.Indent(&out, data, "", "\t")
	out.WriteTo(file)
	file.Close()
}
