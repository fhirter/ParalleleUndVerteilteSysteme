import {AMQPWebSocketClient} from "./amqp-websocket-client.mjs";

export function AmqpClient(callback) {
    const host = "cow-01.rmq2.cloudamqp.com/ws/amqp";
    const vhost = "fggiprgi";
    const username = "fggiprgi";
    const password = "U_tDo0lsDlqKUZgXQ9ywc4K1D3g_j-jo";

    const url = `wss://${host}}`;
    const amqp = new AMQPWebSocketClient(url, `${vhost}`, username, password);

    let channel;
    let queue;

    start().then(() => {
        callback();
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

    return Object.freeze({
        publish,
        subscribe
    })
}