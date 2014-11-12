// The loginController is responsible for handling logins. It connects up
// the websocket service, and sets the JWT token for the user. It then
// broadcasts the authenticated event so other controllers can enable
// themselves.
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

    // login makes a rest call to authenticate the user. If the call is
    // successful then it saves the returned JWT token, and starts the
    // websocket connections, and signals authenticated. On failure
    // it just clears the username and password.
    $scope.login = function() {
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
    };
}
