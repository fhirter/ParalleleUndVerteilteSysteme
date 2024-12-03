# README

## Run server using docker

```shell
DELAY=1000
PORT=8080

docker build -t go-server .
docker run -p 8080:8080 -e REQUEST_DELAY_MS=$DELAY -e PORT=$PORT go-server
```

## Test the connection

```shell
#EVENT='"000a3db6-07f1-4b7d-ba86-00e67173c1b8","shipment.state.set","1","2024-11-23 11:45:30.789000 +00:00","application/json","[""prepare""]","77670","Shipment/62aa4a44-25e6-44cb-a929-bfec93980e0d"'
EVENT="{}"
curl -X POST http://localhost:8080/api/v1/events \
     -H "Content-Type: application/json" \
     -d $EVENT \
     -w "\nTime: %{time_total} seconds\n" \
     -o /dev/null \
     -s
 ```