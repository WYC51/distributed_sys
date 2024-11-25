# restful-api-demo
 In this repository, I've developed a RESTful API server using Golang. Additionally, I've crafted two client applications in both C++ and Python to seamlessly interact with this server.

---
## Dependency
#### Go
```shell
cd path/to/RESTful-API
go mod tidy
```

#### Python
```shell
cd path/to/RESTful-API/client_py
pip install -r requirements.txt
```

#### CPP
1. [cpp-httplib](https://github.com/yhirose/cpp-httplib)
2. [nlohmann/json](https://github.com/nlohmann/json)
```shell
git clone {cpp-httplib} # refer 1.cpp-httplib
git clone {json} # refer 2.nlohmann/json

mv cpp-httplib /path/to/RESTful-API/client_cpp
mv json /path/to/RESTful-API/client_cpp
```

---
## Run
#### Start Server
```shell
go run server.go
```

#### Client (Python)
```shell
python main.py --method {GET/DELETE/PUT/POST} \
--id ?
```

#### Client (CPP)
```shell
g++ main.cpp -o main # compile *.cpp file
./main {GET/DELETE/PUT/POST} {id}
```
