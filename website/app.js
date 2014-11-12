// This file configures our AngularJS application. It sets up
// the global App variable for the creation of the various types
// of services, controllers, filters and directives.
var App = App || {};

App.Constants = angular.module('app.constants', []);
App.Services = angular.module('app.services', []);
App.Controllers = angular.module('app.controllers', []);
App.Filters = angular.module('app.filters', []);
App.Directives = angular.module('app.directives', []);

// Load our module dependencies.
var app = angular.module('myapp', [
    "ui.router", "restangular",
    "ngWebsocket", "angular-jwt",
    "app.constants", "app.services",
    "app.controllers", "app.filters",
    "app.directives"
]);

// appConfig configures our Angular application. In this case we
// are doing the following:
//   1. Configure our application routes
//   2. Configure $http (used by Restangular) to include the
//      JWT Authorization token in the Request header.
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

    // If the route isn't recognized goto /users
    $urlRouterProvider.otherwise("/users");

    // The JWT token is stored in sessionStorage. When our
    // app starts up we explicitly clear the previous token.
    sessionStorage.setItem("token", null);

    // This interceptor will set the Authorization field
    // in the header with the JWT token.
    jwtInterceptorProvider.tokenGetter = function() {
        var token = sessionStorage.getItem("token");
        return token ? token : "";
    };
    $httpProvider.interceptors.push("jwtInterceptor");
}

// appRun allows us to intercept different events while our
// application is running. Here it is used to control access
// to the application by requiring the user to login.
app.run(["$rootScope", "User", "$state", appRun]);
function appRun($rootScope, User, $state) {
    // $stateChangeStart is fired when a route change is starting.
    // Here we check if the user is already authenticatd. If they
    // aren't then we redirect them to the login page.
    $rootScope.$on('$stateChangeStart', function(event, toState, toParams) {
        if (!User.isAuthenticated()) {
            if (toState.url !== "/login") {
                // Cancel whatever route we were going to
                // and instead go to the login page.
                event.preventDefault();
                $state.go("login");
            }
        }
    });
}
