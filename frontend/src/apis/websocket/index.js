let websocket = new WebSocket('ws://' + document.location.hostname + ':9000/ws');

websocket.onopen = function (e) {
    console.log("Connection is established");
};

websocket.onerror = function (error) {
    console.log(`Socket connection error: ${error.code}`);
};

let defineWsListeners = function (processSocketData) {

    websocket.onmessage = function (message) {
        console.log(`Data from WebSocket: ${message}`);
        processSocketData(message);
    };

    websocket.onclose = function (event) {
        if (event.wasClean) {
            console.log(`Clean connection closure, code: ${event.code}, reason: ${event.reason}`);
        } else {
            console.log(`Connection was interrupted, code: ${event.code}`);
        }
    };
};

let sendMessage = function(message) {
    console.log("Message sending...");
    websocket.send(message);
};

export { defineWsListeners, sendMessage }
