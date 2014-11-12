App.Controllers.controller("indexController",
                           ["$scope", "ws",
                            indexController]);

function indexController($scope, ws) {
    $scope.model = {
        connectionStatus: "Not Connected",
        connected: false
    };

    var s = ws.get();

    s.$on("$open", function() {
        console.log("connected");
        $scope.model.connectionStatus = "Connected";
        $scope.model.connected = true;
    });

    s.$on("$close", function() {
        console.log("not connected");
        $scope.model.connectionStatus = "Not Connected";
        $scope.model.connected = false;
    });

    s.$on("$error", function(e) {
        console.log("An error occurred: %O", e);
    });
}
