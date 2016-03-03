/*------------------------------------------------------------------------------
-- DATE:	       March 1, 2016
--
-- Source File:	 check_bans.go
--
-- REVISIONS: 	(Date and Description)
--
-- DESIGNER:	   Marc Vouve
--
-- PROGRAMMER:	 Marc Vouve
--
--
-- INTERFACE:
--	func (m *manifestType) checkBans()
--  func dropBan(ip string)
--
-- NOTES: This file was moved out of main.go
------------------------------------------------------------------------------*/
package main

import (
	"os/exec"
	"time"
)

/*-----------------------------------------------------------------------------
-- FUNCTION:    checkBans
--
-- DATE:        February 25, 2016
--
-- REVISIONS:	  (none)
--
-- DESIGNER:		Marc Vouve
--
-- PROGRAMMER:	Marc Vouve
--
-- INTERFACE:		func (m *manifestType) checkBans()
-- currentBan
--
-- RETURNS: 		void
--
-- NOTES:			This function checks if any of the bans in the ban list should have
--						expired before the script was run. If there are bans that should expired
--						it drops the ban, and removes it from the queue.
------------------------------------------------------------------------------*/
func (m *manifestType) checkBans() {
	for ip, expiry := range m.Bans {
		if time.Now().After(expiry) {
			dropBan(ip)
			delete(m.Bans, ip)
		}
	}
}

/*-----------------------------------------------------------------------------
-- FUNCTION:    dropBan
--
-- DATE:        February 25, 2016
--
-- REVISIONS:	  (none)
--
-- DESIGNER:		Marc Vouve
--
-- PROGRAMMER:	Marc Vouve
--
-- INTERFACE:		func dropBan(ip string)
-- currentBan
--
-- RETURNS: 		void
--
-- NOTES:			This function uses iptables to ban a user from all the basic netfilter
--						chains.
------------------------------------------------------------------------------*/
func dropBan(ip string) {
	for _, chain := range []string{"INPUT", "OUTPUT", "FORWARD"} {
		exec.Command("iptables", "-D", chain, "-s", ip, "-j", "DROP").Run()
	}
}
