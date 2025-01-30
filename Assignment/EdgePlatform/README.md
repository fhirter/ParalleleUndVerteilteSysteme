- `brew install mosquitto`

in 3 separate terminals:

- run broker: `mosquitto`
- run subscriber: `mosquitto_sub -h broker -t 'temp'`
- run publisher: `mosquitto_pub -h localhost -t 'temp' -m "test"`

or using docker:

- `docker-compose up`