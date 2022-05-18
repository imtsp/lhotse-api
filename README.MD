# lhotse 接口协议

## Install
安装buf工具
```shell
brew install bufbuild/buf/buf
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## protobuf gen
```shell
./gen.sh
```

## Java 打包
```shell
mvn clean package -DskipTests
```