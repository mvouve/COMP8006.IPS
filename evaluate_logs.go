package main

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"time"
)

func (m *manifestType) checkSecure() {
	filepath := configure.MustValue("auth", "file", "/var/log/secure")
	logFile, err := os.Open(filepath)
	fileError(filepath, err)
	defer logFile.Close()

	if _, err = logFile.Seek(m.FilePos, os.SEEK_SET); err != nil {
		logFile.Seek(0, 0)
		m.FilePos = 0
	}
	logBuff := bufio.NewReader(logFile)
	for {
		if line, err := logBuff.ReadString('\n'); err == io.EOF {
			break
		} else {
			m.FilePos += int64(len(line))
			if isFailedLogin(line) {
				m.addEvent(getIPfromString(line), getTimeFromString(line))
			}
		}
	}
}

func (m *manifestType) addEvent(ip string, eventTime time.Time) {
	if m.Events[ip] == nil {
		m.Events[ip] = make([]time.Time, 0, 1024)
	}
	m.Events[ip] = append(m.Events[ip][0:], eventTime)
}

func isFailedLogin(log string) bool {
	failed, _ := regexp.MatchString("Failed password", log)

	return failed
}
