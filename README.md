# reverse-sock-over-smb-pipe-named

compile `go build file.go`

*this project is only available on Windows because it use named pipe a Windows technologie*

**Warning**

For named pipe accept remote connection you must modify `%GOPATH%/src/gihub/microsoft/go-winio/pip.go` set `cFILE_PIPE_REJECT_REMOTE_CLIENTS = 8`  to `cFILE_PIPE_REJECT_REMOTE_CLIENTS = 0`

#### Reverse sock
client sock (init pipe connection):
```
main.exe -connect 127.0.0.1
```
server sock (listen pipe connection) :
```
main.exe -listen true -socks 127.0.0.1:8080
```


#### Simple sock
client sock (listen pipe connection) sock:
```
main.exe -listen 127.0.0.1
```
server sock (init pipe connection) :
```
main.exe  -connect 127.0.0.1 -socks 127.0.0.1:8080
```
#### Thanks
This project is hightly inspired by [https://github.com/brimstone/rsocks](https://github.com/brimstone/rsocks) 
