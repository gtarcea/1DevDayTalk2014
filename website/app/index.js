App.Controllers.controller("indexController",
                           ["$scope", "ws",
                            indexController]);

function indexController($scope, ws) {
    $scope.connected = "Not Connected";

    var s = ws.get();

    s.$on("$open", function() {
        $scope.connected = "Connected";
    });

    s.$on("$close", function() {
        $scope.connected = "Not Connected";
    });
}
