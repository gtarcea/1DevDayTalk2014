App.Controllers.controller("usersController", ["$scope", usersController]);
function usersController($scope) {
    $scope.users = [
        {
            name: "Margaret Michaels"
        },

        {
            name: "Bob Biff"
        },

        {
            name: "Hank Hile"
        }
    ];
}

App.Controllers.controller("addUserController", ["$scope", addUserController]);
function addUserController($scope) {

}
