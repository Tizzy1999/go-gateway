## User Manual

### Run API Gateway Backend

- git clone this project

`git clone https://github.com/Tizzy1999/go-gateway`

- make sure the local environment install Go 1.12+

```
go version
go version go1.12.15 darwin/amd64
```

- download dependencies

```
export GO111MODULE=on && export GOPROXY=https://goproxy.cn
cd gateway_demo
go mod tidy
```

- execute `go run main.go` in related folder


- create db and import data

```
mysql -h localhost -u root -p -e "CREATE DATABASE go_gateway DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;"
mysql -h localhost -u root -p go_gateway < go_gateway.sql --default-character-set=utf8
```

- modify mysql、redis configuration files

modify ./conf/dev/mysql.toml and ./conf/dev/redis.toml to the local environment.

- run dashboard、proxy services

start dashboard service
```
go run main.go -config=./conf/dev/ -endpoint dashboard
```

start proxy service
```
go run main.go -config=./conf/dev/ -endpoint server
```

### Run Admin Dashboard

- git clone this project

```
git clone https://github.com/Tizzy1999/go-gateway-view-demo
```

- make sure the local environment has installed nodejs

```
node -v
v11.9.0
```

- install node dependency package

```
cd go_gateway_view
npm install
npm install -g cnpm --registry=https://registry.npm.taobao.org
cnpm install
```

- start frontend project

```
npm run dev
```


## Backend Environment Setup and Editor Tutorial

go environment setup intro
http://docscn.studygolang.com/doc/install

go grammar 
http://tour.studygolang.com/welcome/1

go mod
https://blog.csdn.net/l7l1l0l/article/details/102491573

goland basic usage
https://www.cnblogs.com/happy-king/p/9191356.html


## Frontend Environment Setup and Editor Tutorial

install nodejs 
https://nodejs.org/zh-cn/download/


