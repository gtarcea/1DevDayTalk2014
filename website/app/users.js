App.Controllers.controller("usersController",
                           ["$scope", "Restangular", "$timeout", "ws", usersController]);
function usersController($scope, Restangular, $timeout, ws) {
    var s = ws.get();
    $scope.users = [];

    function getUsers() {
        Restangular.one("api").all("users").getList().then(function(users) {
            users.forEach(function(user) {
                $scope.users.push({email: user.email, fullname: user.fullname});
            });
        });
    }

    $scope.$on("newuser", function(event, user) {
        $timeout(function() {
            //console.log("add user %O", user);
            $scope.users.push(user);
        });
    });

    s.$on("$open", function() {
        getUsers();
    });

    s.$on("$close", function() {
        $scope.users = [];
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
