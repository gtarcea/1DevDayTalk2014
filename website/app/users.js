App.Controllers.controller("usersController",
                           ["$scope", "Restangular", "$timeout", usersController]);
function usersController($scope, Restangular, $timeout) {
    $scope.users = [];
    Restangular.one("api").all("users").getList().then(function(users) {
        users.forEach(function(user) {
            $scope.users.push({email: user.email, fullname: user.fullname});
        });
    });

    $scope.$on("newuser", function(event, user) {
        $timeout(function() {
            $scope.users.push(user);
        });
    });
}

App.Controllers.controller("addUserController",
                           ["$scope", "Restangular", "ws", "$rootScope",
                            addUserController]);
function addUserController($scope, Restangular, ws, $rootScope) {
    var s = ws.get();

    $scope.addUser = function() {
        Restangular.one("api").all("users").post({
            fullname: $scope.username,
            email: $scope.email
        }).then(function(user) {
            s.$emit("adduser", {email: user.email, fullname: user.fullname});
        });
        $scope.username = "";
        $scope.email = "";
    };

    s.$on("adduser", function(user) {
        $rootScope.$broadcast("newuser", user);
    });
}
