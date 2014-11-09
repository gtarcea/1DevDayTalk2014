App.Controllers.controller("indexController",
                           ["$scope", "$timeout", "ws",
                            indexController]);

function indexController($scope, $timeout, ws) {
    console.log("indexController");
    $scope.model = {
        connectionStatus: "Not Connected"
    };

    var s = ws.get();

    s.$on("$open", function() {
        console.log("connected");
        $timeout(function() {
            console.log("setting connection status");
            $scope.model.connectionStatus = "Connected";
        }, 100);
        s.$emit("user", "user message");
    });

    s.$on("$close", function() {
        console.log("not connected");
        $scope.model.connectionStatus = "Not Connected";
    });

    $scope.updateStatus = function() {
        $scope.model.connectionStatus = "Checking";
    };
}
