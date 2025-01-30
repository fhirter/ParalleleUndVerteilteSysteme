FROM python:3.12-bullseye

WORKDIR /app
RUN pip install paho-mqtt

RUN apt-get update && apt-get install -y mosquitto mosquitto-clients

COPY . .
CMD [ "./run.sh" ]
