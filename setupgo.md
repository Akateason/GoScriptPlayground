# 配置go环境
### setup
```
go mod init example/hello
```

### 更新库
```
go mod tidy
```
# 如果timetout 可以试试
# go env -w GOPROXY=https://goproxy.cn,direct

### 运行
```
go run .
```


### go env 
```
export PATH="/Users/teason23/go/bin:$PATH"
```


### 将我们的命令安装到 $GOPATH/bin 目录:
$ go install 
