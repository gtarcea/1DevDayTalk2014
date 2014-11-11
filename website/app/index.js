App.Controllers.controller("indexController",
                           ["$scope", "ws",
                            indexController]);

function indexController($scope, ws) {
    $scope.model = {
        connectionStatus: "Not Connected"
    };

    var s = ws.get();

    s.$on("$open", function() {
        $scope.model.connectionStatus = "Connected";
    });

    s.$on("$close", function() {
        $scope.model.connectionStatus = "Not Connected";
    });

    s.$on("$error", function(e) {
        $timeout(function() {
            console.log("An error occured: %O", e);
        });
    });
}
