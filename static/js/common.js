
var myApp = angular.module('myApp', []);

function HeaderCtrl($scope) {
  $scope.header = {name: "header.html", url: "header.html"}
}

function FooterCtrl($scope) {
  $scope.footer = {name: "footer.html", url: "footer.html"}
}

