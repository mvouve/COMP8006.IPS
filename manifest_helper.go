/*------------------------------------------------------------------------------
-- DATE:	       March 1, 2016
--
-- Source File:	 manifest_helper.go
--
-- REVISIONS:
--
-- DESIGNER:	   Marc Vouve
--
-- PROGRAMMER:	 Marc Vouve
--
--
-- INTERFACE:
--	func initManifest() manifestType
--  func loadManifest(file *os.File) manifestType
--  func (m manifestType) save(f string)
--
--
-- NOTES: This file is the main file in the IPS suite.
------------------------------------------------------------------------------*/
package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"time"
)

/*-----------------------------------------------------------------------------
-- FUNCTION:    initManifest
--
-- DATE:        February 25, 2016
--
-- REVISIONS:	  (none)
--
-- DESIGNER:		Marc Vouve
--
-- PROGRAMMER:	Marc Vouve
--
-- INTERFACE:		func initManifest() manifestType
--
-- RETURNS: 		manifestType a new manifest.
--
-- NOTES:			This function creates a new manifest.
------------------------------------------------------------------------------*/
func initManifest() manifestType {
	events := make(map[string][]time.Time)
	bans := make(map[string]time.Time)

	return manifestType{Events: events, Bans: bans}
}

/*-----------------------------------------------------------------------------
-- FUNCTION:    loadManifest
--
-- DATE:        February 25, 2016
--
-- REVISIONS:	  (none)
--
-- DESIGNER:		Marc Vouve
--
-- PROGRAMMER:	Marc Vouve
--
-- INTERFACE:		func loadManifest(file *os.File) manifestType
--		  file:		a file to load the manifest from.
--
-- RETURNS: 		manifestType - the manifest loaded from the file
--
-- NOTES:			This function loads a manfiest from a file.
------------------------------------------------------------------------------*/
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

/*-----------------------------------------------------------------------------
-- FUNCTION:    save
--
-- DATE:        February 25, 2016
--
-- REVISIONS:	  (none)
--
-- DESIGNER:		Marc Vouve
--
-- PROGRAMMER:	Marc Vouve
--
-- INTERFACE:		func (m manifestType) save(f string)
--				 f:   name of the file to write to, doesn't need to exist.
--
-- RETURNS: 		void
--
-- NOTES:			saves the manifest to a file
------------------------------------------------------------------------------*/
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
