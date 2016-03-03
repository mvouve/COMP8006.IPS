/*------------------------------------------------------------------------------
-- DATE:	       February 25, 2016
--
-- Source File:	 main.go
--
-- REVISIONS: 	March 1, 2016
--									Moved many functions out of this file into their own files.
--
-- DESIGNER:	   Marc Vouve
--
-- PROGRAMMER:	 Marc Vouve
--
--
-- INTERFACE:
--	func main()
--
-- NOTES: This file is the main file in the IPS suite.
------------------------------------------------------------------------------*/
package main

import (
	"os"

	"github.com/Unknwon/goconfig"
)

var configure *goconfig.ConfigFile

/*-----------------------------------------------------------------------------
-- FUNCTION:    main
--
-- DATE:        February 25, 2016
--
-- REVISIONS:	  (none)
--
-- DESIGNER:		Marc Vouve
--
-- PROGRAMMER:	Marc Vouve
--
-- INTERFACE:		func main()
--
-- RETURNS: 		void
--
-- NOTES:			The main entry point for the ips.
------------------------------------------------------------------------------*/
func main() {
	var ips manifestType
	var err error

	configure, err = goconfig.LoadConfigFile(configIni)
	fileError(configIni, err)

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
