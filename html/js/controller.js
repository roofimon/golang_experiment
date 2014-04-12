function MainCtrl($scope, $http) {
$scope.Username = "pallat"
$scope.Password = "go"
  $http.post('http://localhost:8080/post', {
      Username: $scope.Username,
      Password: $scope.Password,
  })
  .success(function(data, status, headers, config) {
      $scope.message = data.Name;
  })
  .error(function(err, status, headers, config) {
      console.log("Well, this is embarassing.");
  });
}