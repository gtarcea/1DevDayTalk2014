// indexController performs book keeping tasks on the index.html including
// showing the websocket connection status.
App.Controllers.controller("indexController",
                           ["$scope", "ws", "$state", "User",
                            indexController]);

function indexController($scope, ws, $state, User) {
    $scope.model = {
        connectionStatus: "Not Connected",
        connected: false,
        authenticated: false
    };

    // Wait on authenticated event. When received then the websocket
    // has been setup and we can start monitoring it.
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

    // logout will clear all login state, close the websocket connection
    // and return us to the login page.
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
