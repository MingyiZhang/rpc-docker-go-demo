# RPC Between Containers

This is a simple example that shows how jsonrpc works between containers in Go.

The example is for running a server which provide a division operation while a 
client send two numbers `a` and `b` to server and get `a/b` as result. 

## How to run
To start everything:
```shell script
docker-compose up --build
```
To stop everything:
```shell script
docker-compose down
```

