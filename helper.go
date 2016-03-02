package main

import (
	"log"
	"os"
	"regexp"
	"time"
)

func fileError(str string, err error) {
	if os.IsNotExist(err) {
		log.Fatalf("%s not found\n", str)
	} else if err != nil {
		log.Fatalln("Error opening %v: %v", str, err)
	}
}

func getIPfromString(log string) string {
	regx := regexp.MustCompile(ipRegex)

	return regx.FindString(log)
}

func getTimeFromString(log string) time.Time {
	t, _ := time.Parse(dateFmt, getTimeStringFromString(log))

	return t
}

func getTimeStringFromString(line string) string {
	regx := regexp.MustCompile(timeStampRegex)

	return regx.FindString(line)
}
