import time

import paho.mqtt.subscribe as subscribe


hostname="localhost"
port=1883

def on_message_print(client, userdata, message):
    print("%s %s" % (message.topic, message.payload))


subscribe.callback(on_message_print, ["temperature", 'humidity', 'on_led'], hostname=hostname, port=port, userdata={"message_count": 0})

while True:
    time.sleep(1)