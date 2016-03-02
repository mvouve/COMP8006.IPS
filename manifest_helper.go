package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"time"
)

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

func (m manifestType) save(f string) {
	data, err := json.Marshal(m)
	if err != nil {
		log.Fatalln("Error mashaling JSON: ", err)
	}
	file, _ := os.Create(f)
	defer file.Close()

	var out bytes.Buffer
	json.Indent(&out, data, "", "\t")
	out.WriteTo(file)
	file.Close()
}
