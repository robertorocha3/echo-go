# Purpose

This is an echo service used for testing purposes.

# Pre-requisites

* Go
* jq (optional but recommended)

## Working locally

### How to compile then run it

```
go build -o app . && ./app
```

### How to call it

```
curl -s -X POST -d '{
   "Request": "repeat this"
}' "http://localhost:3000/echo" | jq .
```

### Sample response

```
{
  "description": "echo-go",
  "result": "SUCCESS",
  "value": "repeat this"
}
```

## Working with Kubernetes

### How to build it

```shell script
./run.sh build
```

### How to deploy it

```shell script
./run.sh deploy
```

### To build then deploy

```shell script
./run.sh build-and-deploy
```

### How to call it after deployment

```shell script
curl -s -X POST -d '{
   "Request": "repeat that"
}' "http://echo-go.dev-simple-apps.tigersoftware.local:10080/echo" | jq .
```

### Sample response

```
{
  "description": "echo-go",
  "result": "SUCCESS",
  "value": "repeat that"
}
```

## Other endpoints

### Get the Prometheus metrics

```shell script
curl -s -X GET "http://localhost:3000/metrics"
``` 

### Get this app's info

```shell script
curl -s -X GET "http://localhost:3000/info" | jq .
``` 
