# gRPC_Server_Client
A straightforward server-client model built on gRPC

---
### Environment
- go 1.21.6

### Installation
```shell
# install all dependencies
go mod tidy 
```

### Run Server

```shell
go run server/server.go
```

### Run Client
```shell
go run client/client.go
```

---
## Demo

### Client Side
```shell
# run command on terminal
go run client/client.go

#output
2024/03/09 08:56:05 Greeting: Server say hello to {NAME}
```

### Server Side
```shell
# run command on terminal
go run server/server.go

#output
starting gRPC server on port :{PORT NUM}
2024/03/09 08:56:05 Received: {NAME}
```
