package main

import "time"

const configIni string = "/root/go/src/github.com/mvouve/COMP8006.IPS/config.ini"
const defaultTraceTime int = 1
const defaultTrys int = 5 // default ammount of tries allowed (overridden by config)
const defaultBan int = 24 // time in hours for default ban (overriden by config)
const dateFmt string = "Jan _2 15:04:05"
const timeStampRegex string = `[A-Za-z]{3}\s+\d{1,2}\s\d{2}:\d{2}:\d{2}`
const ipRegex string = `[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}`

type manifestType struct {
	FilePos int64
	Bans    map[string]time.Time
	Events  eventList
}

type eventList map[string][]time.Time
