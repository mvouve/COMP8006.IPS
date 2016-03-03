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
--	func (m *manifestType) checkEvents()
--  func (m *manifestType) addBan(ip string)
--
-- NOTES: This file was moved out of main.go
------------------------------------------------------------------------------*/
package main

import (
	"os/exec"
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
-- INTERFACE:		func (m *manifestType) checkEvents()
--
--
-- RETURNS: 		void
--
-- NOTES:			This function checks if any of the bans in the ban list should have
--						expired before the script was run. If there are bans that should expired
--						it drops the ban, and removes it from the queue.
------------------------------------------------------------------------------*/
func (m *manifestType) checkEvents() {
	traceTime := time.Minute * time.Duration(configure.MustInt("auth", "trace_time", defaultTraceTime))
	now, _ := time.Parse(dateFmt, time.Now().Format(dateFmt))
	for ip := range m.Events {
		recentEvents := make([]time.Time, 0, 128)
		for _, instance := range m.Events[ip] {
			if now.Sub(instance) < traceTime {
				recentEvents = append(recentEvents[0:], instance)
			}
		}
		if len(recentEvents) >= configure.MustInt("auth", "max_attempts", defaultTrys) {
			m.addBan(ip)
		}
	}
}

/*-----------------------------------------------------------------------------
-- FUNCTION:    addBan
--
-- DATE:        February 25, 2016
--
-- REVISIONS:	  (none)
--
-- DESIGNER:		Marc Vouve
--
-- PROGRAMMER:	Marc Vouve
--
-- INTERFACE:		func (m *manifestType) addBan(ip string)
-- 				ip: 	IP to add to netfilter and m's ban list.
--
-- RETURNS: 		void
--
-- NOTES:			This function adds a ip to the ban tracker and netfilter
------------------------------------------------------------------------------*/
func (m *manifestType) addBan(ip string) {
	_, ok := m.Bans[ip]
	if !ok {
		for _, chain := range []string{"INPUT", "OUTPUT", "FORWARD"} {
			exec.Command("iptables", "-A", chain, "-s", ip, "-j", "DROP").Run()
		}
	}
	m.Bans[ip] = time.Now().Add(time.Hour * time.Duration(configure.MustInt("auth", "ban_time", defaultBan)))
}
