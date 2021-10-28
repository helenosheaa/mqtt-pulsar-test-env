MQTT Pulsar Test Env

Setup apache pulsar:

`docker-compose up`

If Docker-compose fails to start up successfully run:

```
docker run -it \
-p 6650:6650 \
-p 8080:8080 \
-v $PWD/data:/pulsar/data \
  apachepulsar/pulsar:2.4.0 \
  bin/pulsar standalone
```

Setup pulsar manager:

Set token:

`CSRF_TOKEN=$(curl http://localhost:7750/pulsar-manager/csrf-token)`

Set user:

```
curl \
    -H "X-XSRF-TOKEN: $CSRF_TOKEN" \
    -H "Cookie: XSRF-TOKEN=$CSRF_TOKEN;" \
    -H 'Content-Type: application/json' \
    -X PUT http://localhost:7750/pulsar-manager/users/superuser \
    -d '{"name": "admin", "password": "apachepulsar", "description": "test", "email": "username@test.org"}'
 ```

Navigate to manager UI:

`http://localhost:9527/`

Setup new environment:

Service URL:

`http://pulsar:8080`

Start pulsar client, consumer and producer: 

`go build main.go`
`./main`

View topics in manager UI

Add MQTT MOP:

