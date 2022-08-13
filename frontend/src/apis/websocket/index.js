let websocket = new WebSocket('ws://localhost:9000/ws');

let connect = function (processSocketData) {
    
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
        console.log(`Socket connection error: ${error.code}`);
    };
};

let sendMessage = function(message) {
    console.log("Message sending...");
    websocket.send(message);
};

export { connect, sendMessage }
