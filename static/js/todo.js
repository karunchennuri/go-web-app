/*
 Copyright 2011 The Go Authors.  All rights reserved.
 Use of this source code is governed by a BSD-style
 license that can be found in the LICENSE file.
*/

function TaskCtrl($scope, $http, $window) {
  $scope.tasks = [];
  $scope.working = false;

  var logError = function(data, status) {
    console.log('code '+status+': '+data);
    $scope.working = false;
  };

  var refresh = function() {
    return $http.get('/task').
      success(function(data) { $scope.tasks = data.Tasks; }).
      error(logError);
  };

  $scope.addTodo = function() {
    $scope.working = true;
    $http.post('/task', {Title: $scope.todoText}).
      error(logError).
      success(function() {
        refresh().then(function() {
          $scope.working = false;
          $scope.todoText = '';
        })
      });
  };

  $scope.toggleDone = function(task) {
    data = {ID: task.ID, Title: task.Title, Done: !task.Done}
    $http.put('/task/'+task.ID, data).
      error(logError).
      success(function() { task.Done = !task.Done });
  };

  $scope.logout = function() {
      $scope.working = true;
      $http.get('/', {token: $scope.token}).
        error(logError).
        success(function() {
          refresh().then(function() {
            $scope.working = false;
            $scope.tasks = [];
            $window.location.href='/';
          })
        });
    };

  refresh().then(function() { $scope.working = false; });
}


function LogIn($scope, $http, $location, $window) {
    $scope.token = '';
    $scope.tasks = [];
    $scope.working = false;
    var logError = function(data, status) {
      console.log('code '+status+': '+data);
      alert('Invalid token! Try again')
      $scope.working = false;
    };

    $scope.login = function() {
      $scope.working = true;
      //alert('token=' + $scope.token)
      $http.post('/login', {Token: $scope.tokenText}).
        error(logError).
        success(function() {
          alert("SUCCESS login");
          $scope.tasks = [];
          $window.location.href='/test';
        });
    }

    $scope.toggleDone = function(token) {
      data = {token: token, Done: !login.Done}
    };

    $scope.autoExpand = function(e) {
            var element = typeof e === 'object' ? e.target : document.getElementById(e);
        		var scrollHeight = element.scrollHeight -60; // replace 60 by the sum of padding-top and padding-bottom
            element.style.height =  scrollHeight + "px";
        };

      function expand() {
        $scope.autoExpand('TextArea');
      }
}
