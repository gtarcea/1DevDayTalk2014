App.Controllers.controller("indexController",
                           ["$scope", "ws", "$state", "User",
                            indexController]);

function indexController($scope, ws, $state, User) {
    $scope.model = {
        connectionStatus: "Not Connected",
        connected: false,
        authenticated: false
    };

    $scope.$on("authenticated", function() {
        $scope.model.authenticated = true;
        var s = ws.get();

        s.$on("$open", function() {
            $scope.model.connectionStatus = "Connected";
            $scope.model.connected = true;
        });

        s.$on("$close", function() {
            $scope.model.connectionStatus = "Not Connected";
            $scope.model.connected = false;
        });

        s.$on("$error", function(e) {
            console.log("An error occurred: %O", e);
        });
    });

    $scope.logout = function() {
        $scope.model.authenticated = false;
        $scope.model.connected = false;
        $scope.model.connectionStatus = "Not Connected";
        var s = ws.get();
        s.$close();
        User.setToken(null);
        User.setAuthenticated(false);
        $state.go("login");
    };
}
