var App = App || {};

App.Constants = angular.module('app.constants', []);
App.Services = angular.module('app.services', []);
App.Controllers = angular.module('app.controllers', []);
App.Filters = angular.module('app.filters', []);
App.Directives = angular.module('app.directives', []);

var app = angular.module('myapp', [
    "ui.router", "restangular",
    "app.constants", "app.services",
    "app.controllers", "app.filters",
    "app.directives"
]);

app.config(["$stateProvider", "$urlRouterProvider", appConfig]);
function appConfig($stateProvider, $urlRouterProvider) {
    $stateProvider
        .state("users", {
            url: "/users",
            templateUrl: "app/users.html",
            controller: "usersController"
        })
        .state("users.add", {
            url: "/add",
            templateUrl: "app/add.html",
            controller: "addUserController"
        });
        $urlRouterProvider.otherwise("/users");
}

app.run([appRun]);
function appRun() {

}
