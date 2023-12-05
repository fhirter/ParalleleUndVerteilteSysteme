import {AMQPWebSocketClient} from './amqp-websocket-client.mjs'

const host = "cow-01.rmq2.cloudamqp.com/ws/amqp";
const vhost = "";
const username = "";
const password = "";

const url = `wss://${host}}`;
const amqp = new AMQPWebSocketClient(url, `${vhost}`, username, password);

const output = document.getElementById("output");
const input = document.getElementById("message");

async function start() {
    try {
        const connection = await amqp.connect()
        const channel = await connection.channel()
        attachPublish(channel)
        const queue = await channel.queue("")
        await queue.bind("amq.fanout")
        const consumer = await queue.subscribe({noAck: false}, (msg) => {
            console.log("from broker:", msg)
            output.innerText += msg.bodyToString() + "\n"
            msg.ack()
        })
    } catch (err) {
        console.error("Error", err, "reconnecting in 1s")
        disablePublish()
        setTimeout(start, 1000)
    }
}

function attachPublish(ch) {
    document.forms[0].onsubmit = async (e) => {
        e.preventDefault()
        try {
            await ch.basicPublish("amq.fanout", "", input.value, {contentType: "text/plain"})
        } catch (err) {
            console.error("Error", err, "reconnecting in 1s")
            disablePublish()
            setTimeout(start, 1000)
        }
        input.value = ""
    }
}

function disablePublish() {
    document.forms[0].onsubmit = (e) => {
        alert("Disconnected, waiting to be reconnected")
    }
}

start()