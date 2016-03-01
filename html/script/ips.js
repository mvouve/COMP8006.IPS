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

    hostList.currentBan = function(host) {
      var t;
      angular.forEach(hostList.manifest.Bans, function(value, key) {
        if(key === host) {
          t = value;
        }
      });
      return t
    };

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
