export function MqttClient(config, callback) {
    const client = mqtt.connect(config.host);

    const channel = "chat";

    const publicClient = Object.freeze({
        publish,
        subscribe
    })

    client.on("connect", () => {
        client.subscribe(channel, (err) => {
            if (!err) {
                callback(publicClient);
            } else {
                console.error(err);
            }
        })
    });

    function publish(message) {
        client.publish(channel, message);
    }

    function subscribe(callback) {
        client.on("message", (topic, message) => {
            // message is Buffer
            console.log(message.toString());
            callback(message.toString());
        });
    }


}
