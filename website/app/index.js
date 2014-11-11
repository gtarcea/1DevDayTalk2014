App.Controllers.controller("indexController",
                           ["$scope", "$timeout", "ws",
                            indexController]);

function indexController($scope, $timeout, ws) {
    $scope.model = {
        connectionStatus: "Not Connected"
    };

    var s = ws.get();

    s.$on("$open", function() {
        $timeout(function() {
            $scope.model.connectionStatus = "Connected";
        });
    });

    s.$on("$close", function() {
        $timeout(function() {
            $scope.model.connectionStatus = "Not Connected";
        });
    });

    s.$on("$error", function(e) {
        $timeout(function() {
            console.log("An error occured: %O", e);
        });
    });
}
