import {AmqpClient} from "./amqpClient.mjs";
import {config} from "./config.mjs";
import {MqttClient} from "./mqttClient.mjs";

const output = document.getElementById("output");
const input = document.getElementById("message");

const callback = (client) => {
    // publish messages
    document.forms[0].onsubmit = async (e) => {
        e.preventDefault()
        await client.publish(input.value);
        input.value = ""
    }

    // subscribe
    client.subscribe((message) => {
        output.innerText += message + "\n";
    })
}

// uncomment client that should be used
const amqpClient = new AmqpClient(config.amqp, callback);
// const mqttClient = new MqttClient(config.mqtt, callback);
