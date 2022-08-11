let websocket;

let connect = function (processSocketData) {
    websocket = new WebSocket("wss://localhost:9090/ws");

    websocket.onopen = function (e) {
        console.log("Connection is established");
    };

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

    websocket.onerror = function (error) {
        console.log(`Socket connection error: ${error}`);
    };
};

let sendMessage = function(message) {
    console.log("Message sending...");
    websocket.send(message);
};

export { connect, sendMessage }
