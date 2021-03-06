From golang
MAINTAINER "ZZW"
RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct && go get -u gorm.io/gorm && go get -u gorm.io/driver/sqlite
COPY . .
RUN cd src/hello && go mod init hello && go mod tidy && go build main.go
