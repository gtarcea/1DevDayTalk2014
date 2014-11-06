App.Controllers.controller("usersController", ["$scope", "Restangular", usersController]);
function usersController($scope, Restangular) {
    Restangular.one("api").all("users").getList().then(function(users) {
        $scope.users = users;
    });
}

App.Controllers.controller("addUserController", ["$scope", "Restangular", addUserController]);
function addUserController($scope, Restangular) {
    $scope.addUser = function() {
        Restangular.one("api").all("users").post({
            fullname: $scope.username,
            email: $scope.email
        }).then(function(user) {
            $scope.users.push(user);
        });
        $scope.username = "";
        $scope.email = "";
    };
}
