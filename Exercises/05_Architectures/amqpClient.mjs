import {AMQPWebSocketClient} from "./amqp-websocket-client.mjs";

export function AmqpClient(config, callback) {
    const url = `wss://${config.host}}`;
    const amqp = new AMQPWebSocketClient(url, `${config.vhost}`, config.username, config.password);

    let channel;
    let queue;

    const publicClient = Object.freeze({
        publish,
        subscribe
    })

    start().then(() => {
        callback(publicClient);
    });

    async function start() {
        try {
            const connection = await amqp.connect()
            channel = await connection.channel()
            queue = await channel.queue("")
            await queue.bind("amq.fanout")
        } catch (err) {
            console.error("Error", err, "reconnecting in 1s")
            setTimeout(start, 1000)
        }
    }

    async function publish(message) {
        try {
            await channel.basicPublish("amq.fanout", "", message, {contentType: "text/plain"})
        } catch (err) {
            console.error("Error", err, "reconnecting in 1s")
            setTimeout(start, 1000)
        }
    }

    function subscribe(callback) {
        queue.subscribe({noAck: false}, (msg) => {
            console.log("from broker:", msg);
            callback(msg.bodyToString());
            msg.ack();
        })
    }
}