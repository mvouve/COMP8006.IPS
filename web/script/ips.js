angular.module('ipsApp', [])
  .controller('HostListController', function($scope, $http) {

    var hostList = this;
    hostList.ipsInfo = {
      	"FilePos": 3944,
      	"Bans": {
      		"10.64.205.150": "2016-02-29T01:59:44.772948802-08:00"
      	},
      	"Events": {
      		"10.64.205.150": [
      			"0000-02-28T01:57:08Z",
      			"0000-02-28T01:57:12Z",
      			"0000-02-28T01:57:15Z",
      			"0000-02-28T01:59:33Z",
      			"0000-02-28T01:59:37Z",
      			"0000-02-28T01:59:42Z"
      		],
          "10.64.205.199": [
      			"0000-02-20T01:57:08Z",
      			"0000-02-20T01:57:12Z",
      			"0000-02-20T01:57:15Z",
      		]
      	}
      }

    hostList.currentBan = function(host) {
      var t
      angular.forEach(hostList.ipsInfo.Bans, function(value, key) {
        console.log(key + " " + host)
        if(key === host) {
          t = value
        }
      })
      return t
    }

    hostList.events = function(host) {
      var e = []

      angular.forEach(hostList.ipsInfo.Events, function(value, key) {
        if(key === host) {
          e = value
        }
      }, e)

      return e;
    }
    hostList.hosts = function() {
      var h = [];
      angular.forEach(hostList.ipsInfo.Events, function(value, key) {
        this.push(key)
      }, h);

      return h;
    }
    var currentHost = hostList.hosts()[0];
  });
