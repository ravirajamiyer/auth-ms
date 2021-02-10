# auth-ms

A microservice for identity, authentication, and authorization

WARNING: This needs a Docker container of postgres DB to try it out. May not work out of the box


## Usage

```
go run main.go
```

## Build Docker image

The below command will build a Docker image with the name 'api'

```
docker build -t api .

``` 


## Run  Docker container

The below command will run the container

```
docker run -it -p 8080:8080 api
``` 

