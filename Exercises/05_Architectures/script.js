import {AmqpClient} from "./amqpClient.mjs";
import {config} from "./config.mjs";

const output = document.getElementById("output");
const input = document.getElementById("message");

const client = new AmqpClient(config, () => {
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
});


