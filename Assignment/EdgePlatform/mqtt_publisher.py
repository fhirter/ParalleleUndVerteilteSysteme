import random

import paho.mqtt.publish as publish

hostname="localhost"
port=1883

temperature = random.random()*50+20
humidity = random.random()*50+20
on_led = random.random()

publish.single("temperature", temperature, hostname=hostname, port=port)
publish.single("humidity", "30", hostname=hostname, port=port)
publish.single("on_led", "30", hostname=hostname, port=port)


