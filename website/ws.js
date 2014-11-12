// The ws service gives us easy access to the websocket.
App.Services.factory("ws", ["$location", "$websocket", wsService]);
function wsService($location, $websocket) {
    var service = {
        // url returns the url to use for the websocket. It tries to be smart
        // to determine the host, port and whether or not it is a secure connection.
        url: function() {
            var protocol = $location.protocol();
            var url = (protocol == "http" ? "ws://" : "wss://") + $location.host() + ":" + $location.port() + "/ws";
            return url;
        },

        // get returns the already connected websocket.
        get: function() {
            var url = service.url();
            return $websocket.$get(url);
        }
    };

    return service;
}
