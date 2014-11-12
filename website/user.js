// The User service tracks the user state. If you want your user state
// to service browser refresh then store the state in a cookie, session
// or local storage depending on your apps needs.
//
// This service persists the JWT to session storage. This provides an
// easy communication mechanism to providers configured in app.config.
// You can only inject provider services into app.config (which this
// service clearly isn't).
App.Services.factory("User", [userService]);

function userService() {
    var self = this;
    self.authenticated = false;
    self.token = "";

    var service = {
        isAuthenticated: function() {
            return self.authenticated;
        },

        setAuthenticated: function(val) {
            self.authenticated = val;
        },

        setToken: function(token) {
            self.token = token;
            sessionStorage.setItem("token", token);
        },

        token: function() {
            return self.token;
        }
    };

    return service;
}
