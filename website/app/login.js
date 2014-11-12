App.Controllers.controller("loginController",
                           ["$scope", "User", "Restangular", "$state", "$rootScope",
                            "$websocket", "ws",
                            loginController]);

function loginController($scope, User, Restangular,
                         $state, $rootScope, $websocket, ws) {
    $scope.cancel = function() {
        $scope.username = "";
        $scope.password = "";
    };

    $scope.login = function() {
        //if ($scope.username == "admin" && $scope.password == "abc123") {
        Restangular.one("api").one("users").all("login").post({
            username: $scope.username,
            password: $scope.password
        }).then(function(auth) {
            User.setAuthenticated(true);
            User.setToken(auth.token);
            $websocket.$new({
                url: ws.url(),
                reconnect: true,
                reconnectInterval: 500
            });
            $rootScope.$broadcast("authenticated");
            $state.go("users");
        }, function() {
            $scope.username = "";
            $scope.password = "";
        });
        //}
    };
}
