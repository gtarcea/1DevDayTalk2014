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
