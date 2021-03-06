From golang
MAINTAINER "ZZW"
COPY . .
RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct && cd src/consul_demo && go mod init consul_demo && go mod tidy && go build main.go
