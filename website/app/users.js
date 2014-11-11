/* The usersController displays the list of users. It is also responsible for tracking
 * the state of the websocket connection. If the connection goes down then we clear
 * the list of users. If the connections comes back up then we requery for the current
 * list of users.
 */
App.Controllers.controller("usersController",
                           ["$scope", "Restangular", "$timeout", "ws", usersController]);
function usersController($scope, Restangular, $timeout, ws) {
    var s = ws.get();
    $scope.users = [];

    // getUsers asks the server for all the known users.
    function getUsers() {
        Restangular.one("api").all("users").getList().then(function(users) {
            users.forEach(function(user) {
                $scope.users.push({email: user.email, fullname: user.fullname});
            });
        });
    }

    // When a new user is added it is broadcast on the websocket.
    // We wrap in $timeout to cause the digest to update the view of users.
    s.$on("addeduser", function(user) {
        $timeout(function() {
            $scope.users.push(user);
        });
    });

    // ng-websocket will send a $open event when a connection is
    // established. We ask the server for its list of users when
    // this event occurs.
    s.$on("$open", function() {
        $timeout(function(){
            getUsers();
        });
    });

    // ng-websocket will send a $close when the connection is
    // closed. Clear the list of users when this occurs.
    s.$on("$close", function() {
        $timeout(function(){
            $scope.users = [];
        });
    });
}

// addUserController adds new users. The process for adding a user works
// as follows: When the server has responded with the added user we fire
// off a "adduser" event on the websocket. The server will then broadcast
// an "addeduser" event to all connected clients.
App.Controllers.controller("addUserController",
                           ["$scope", "Restangular", "ws",
                            addUserController]);
function addUserController($scope, Restangular, ws) {
    var s = ws.get();

    // addUser sends a REST POST to create the user. On success it
    // gets the user back and then sends a adduser event on the
    // websocket. The server will then broadcast a addeduser event
    // to all connected clients.
    //
    // Note: The REST call should tell the websocket to broadcast
    // rather than doing a double hop. The hop is in here to demonstrate
    // sending and receiving from the websocket.
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
}
