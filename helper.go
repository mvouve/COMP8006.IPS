/*------------------------------------------------------------------------------
-- DATE:	       March 1, 2016
--
-- Source File:	 helper.go
--
-- REVISIONS: 	(Date and Description)
--
-- DESIGNER:	   Marc Vouve
--
-- PROGRAMMER:	 Marc Vouve
--
--
-- INTERFACE:
--	func fileError(str string, err error)
--  func getIPfromString(log string) string
--  func getTimeFromString(log string) time.Time
--  func getTimeStringFromString(line string) string
--
-- NOTES: This file was moved out of main.go
------------------------------------------------------------------------------*/
package main

import (
	"log"
	"os"
	"regexp"
	"time"
)

/*-----------------------------------------------------------------------------
-- FUNCTION:    fileError
--
-- DATE:        February 25, 2016
--
-- REVISIONS:	  (none)
--
-- DESIGNER:		Marc Vouve
--
-- PROGRAMMER:	Marc Vouve
--
-- INTERFACE:		func fileError(str string, err error)
--			 str:		name of the file
--			 err:		error returned by opening the file
--
-- RETURNS: 		void
--
-- NOTES:			This function checks the secure file for failed login attempts.
------------------------------------------------------------------------------*/
func fileError(str string, err error) {
	if os.IsNotExist(err) {
		log.Fatalf("%s not found\n", str)
	} else if err != nil {
		log.Fatalln("Error opening %v: %v", str, err)
	}
}

func ferror(str string, err error) {
	if err != nil {
		log.Fatalln("error at %v: %v", str, err)
	}
}

/*-----------------------------------------------------------------------------
-- FUNCTION:    getIPfromString
--
-- DATE:        February 25, 2016
--
-- REVISIONS:	  (none)
--
-- DESIGNER:		Marc Vouve
--
-- PROGRAMMER:	Marc Vouve
--
-- INTERFACE:		func getIPfromString(log string) string
--			 log:		a line from a log file to find extract an IP string from.
--
-- RETURNS: 		the ip expressed as a string.
--
-- NOTES:			This function find an IP address in a string.
------------------------------------------------------------------------------*/
func getIPfromString(log string) string {
	regx := regexp.MustCompile(ipRegex)

	return regx.FindString(log)
}

/*-----------------------------------------------------------------------------
-- FUNCTION:    getTimeFromString
--
-- DATE:        February 25, 2016
--
-- REVISIONS:	  (none)
--
-- DESIGNER:		Marc Vouve
--
-- PROGRAMMER:	Marc Vouve
--
-- INTERFACE:		getTimeFromString(log string) time.Time
--			 log:		a line from a log file to find extract an IP string from.
--
-- RETURNS: 		time.Time the time expressed in the log file based on dateFmt
--
-- NOTES:			This function find an IP address in a string.
------------------------------------------------------------------------------*/
func getTimeFromString(log string) time.Time {
	t, _ := time.Parse(dateFmt, getTimeStringFromString(log))

	return t
}

/*-----------------------------------------------------------------------------
-- FUNCTION:    getTimeStringFromString
--
-- DATE:        February 25, 2016
--
-- REVISIONS:	  (none)
--
-- DESIGNER:		Marc Vouve
--
-- PROGRAMMER:	Marc Vouve
--
-- INTERFACE:		getTimeStringFromString(line string) string
--			line: 	the line from the log file which contains a time.
--
-- RETURNS: 		string the time expressed in the string.
--
-- NOTES:			This function find an IP address in a string.
------------------------------------------------------------------------------*/
func getTimeStringFromString(line string) string {
	regx := regexp.MustCompile(timeStampRegex)

	return regx.FindString(line)
}
