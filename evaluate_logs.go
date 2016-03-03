/*------------------------------------------------------------------------------
-- DATE:	       March 1, 2016
--
-- Source File:	 evaluate_logs.go
--
-- REVISIONS: 	(Date and Description)
--
-- DESIGNER:	   Marc Vouve
--
-- PROGRAMMER:	 Marc Vouve
--
--
-- INTERFACE:
--	func (m *manifestType) checkSecure()
--  func (m *manifestType) addEvent(ip string, eventTime time.Time)
--  func isFailedLogin(log string) bool
--
--
-- NOTES: This file was moved out of main.go
------------------------------------------------------------------------------*/

package main

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"time"
)

/*-----------------------------------------------------------------------------
-- FUNCTION:    checkEvents
--
-- DATE:        February 25, 2016
--
-- REVISIONS:	  (none)
--
-- DESIGNER:		Marc Vouve
--
-- PROGRAMMER:	Marc Vouve
--
-- INTERFACE:		func (m *manifestType) checkSecure()
--
--
-- RETURNS: 		void
--
-- NOTES:			This function checks the secure file for failed login attempts.
------------------------------------------------------------------------------*/
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

/*-----------------------------------------------------------------------------
-- FUNCTION:    addEvent
--
-- DATE:        February 25, 2016
--
-- REVISIONS:	  (none)
--
-- DESIGNER:		Marc Vouve
--
-- PROGRAMMER:	Marc Vouve
--
-- INTERFACE:		func (m *manifestType) addEvent(ip string, eventTime time.Time)
--				ip: 	ip of the event
-- eventTime:   when the event occured.
--
--
-- RETURNS: 		void
--
-- NOTES:			Add an event to the manifest.
------------------------------------------------------------------------------*/
func (m *manifestType) addEvent(ip string, eventTime time.Time) {
	if m.Events[ip] == nil {
		m.Events[ip] = make([]time.Time, 0, 1024)
	}
	m.Events[ip] = append(m.Events[ip][0:], eventTime)
}

/*-----------------------------------------------------------------------------
-- FUNCTION:    addEvent
--
-- DATE:        February 25, 2016
--
-- REVISIONS:	  (none)
--
-- DESIGNER:		Marc Vouve
--
-- PROGRAMMER:	Marc Vouve
--
-- INTERFACE:		func isFailedLogin(log string) bool
--			 log:		a string to search for Failed password.
--
--
-- RETURNS: 		bool true if "Failed Password" was found in the string.
--
-- NOTES:			checks a line for the words "Failed Password"
------------------------------------------------------------------------------*/
func isFailedLogin(log string) bool {
	failed, _ := regexp.MatchString("Failed password", log)

	return failed
}
