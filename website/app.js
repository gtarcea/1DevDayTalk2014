var App = App || {};

App.Constants = angular.module('app.constants', []);
App.Services = angular.module('app.services', []);
App.Controllers = angular.module('app.controllers', []);
App.Filters = angular.module('app.filters', []);
App.Directives = angular.module('app.directives', []);

var app = angular.module('myapp', [
    "ui.router", "restangular",
    "ngWebsocket", "angular-jwt",
    "app.constants", "app.services",
    "app.controllers", "app.filters",
    "app.directives"
]);

app.config(["$stateProvider", "$urlRouterProvider", "$httpProvider",
            "jwtInterceptorProvider",
            appConfig]);
function appConfig($stateProvider, $urlRouterProvider, $httpProvider,
                  jwtInterceptorProvider) {
    $stateProvider
        .state("login", {
            url: "/login",
            templateUrl: "app/login.html",
            controller: "loginController"
        })
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

    sessionStorage.setItem("token", null);

    jwtInterceptorProvider.tokenGetter = function() {
        var token = sessionStorage.getItem("token");
        return token ? token : "";
    };
    $httpProvider.interceptors.push("jwtInterceptor");
}

app.run(["$rootScope", "User", "$state", appRun]);
function appRun($rootScope, User, $state) {
    $rootScope.$on('$stateChangeStart', function(event, toState, toParams) {
        if (!User.isAuthenticated()) {
            if (toState.url !== "/login") {
                event.preventDefault();
                $state.go("login");
            }
        }
    });
}
