/*------------------------------------------------------------------------------
-- DATE:	       February 28, 2016
--
-- Source File:	 ips.js
--
-- REVISIONS:
--
-- DESIGNER:	   Marc Vouve
--
-- PROGRAMMER:	 Marc Vouve
--
--
-- INTERFACE:
--	hostList.currentBan = function(host)
--  hostList.events = function(host)
--  hostList.hosts = function()
--
--
-- NOTES: This is the main angular logic file.
------------------------------------------------------------------------------*/
angular.module('ipsApp', [])
  .controller('HostListController', function($scope, $http, $interval) {

    var hostList = this;
    hostList.manifest = {};
    $http.get('/manifest.json?time=' + Date.now()).then(
      function(response){
        hostList.manifest = response.data;
      });

    $interval(function() {
      $http.post('/manifest.json', "").then(
        function(response){
          hostList.manifest = response.data;
        });
    }, 1000);

/*-----------------------------------------------------------------------------
-- FUNCTION:    currentBan
--
-- DATE:        February 25, 2016
--
-- REVISIONS:	  (none)
--
-- DESIGNER:		Marc Vouve
--
-- PROGRAMMER:	Marc Vouve
--
-- INTERFACE:		hostList.currentBan = function(host)
--      host:   the host to check against.
--
-- RETURNS: 		the time that the ban will end (if the host is banned)
--
-- NOTES:			  if no ban is found "nothing" is returned.
------------------------------------------------------------------------------*/
    hostList.currentBan = function(host) {
      var t;
      angular.forEach(hostList.manifest.Bans, function(value, key) {
        if(key === host) {
          t = value;
        }
      });
      return t
    };

/*-----------------------------------------------------------------------------
-- FUNCTION:    currentBan
--
-- DATE:        February 25, 2016
--
-- REVISIONS:	  (none)
--
-- DESIGNER:		Marc Vouve
--
-- PROGRAMMER:	Marc Vouve
--
-- INTERFACE:		hostList.events = function(host)
--      host:   the host to check against.
--
-- RETURNS: 		recent events from the host.
--
-- NOTES:			  if no events is found "nothing" is returned.
------------------------------------------------------------------------------*/
    hostList.events = function(host) {
      var e = []

      angular.forEach(hostList.manifest.Events, function(value, key) {
        if(key === host) {
          e = value;
        }
      }, e);

      return e;
    };

    hostList.hosts = function() {
      var h = [];
      angular.forEach(hostList.manifest.Events, function(value, key) {
        this.push(key);
      }, h);

      return h;
    };
  });
