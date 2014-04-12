function MainCtrl($scope, $http) {
$scope.Username = "pallat"
$scope.Password = "golang"
$http({
    url: 'http://localhost:8080/post',
    method: 'POST',
    headers: {
      'Accept': 'application/json, text/javascript', 
      'Content-Type': 'application/json; charset=utf-8'
    },
    data: {"username": "test","password": "golang"},
}).success(function(data){ $scope.message = data });

}
