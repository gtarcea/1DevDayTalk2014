App.Services.factory("ws", ["$location", "$websocket", wsService]);
function wsService($location, $websocket) {
    var service = {
        url: function() {
            var protocol = $location.protocol();
            var url = (protocol == "http" ? "ws://" : "wss://") + $location.host() + ":" + $location.port() + "/ws";
            console.log(url);
            return url;
        },

        get: function() {
            var url = service.url();
            return $websocket.$get(url);
        }
    };

    return service;
}
