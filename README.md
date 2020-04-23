# reverse-sock-over-smb-pipe-named

compile `go build file.go`


## Warning

For named pipe accept remote connection you must modify `GOPATH/src/gihub/microsoft/go-winio/pip.go` :

```
	cFILE_PIPE_REJECT_REMOTE_CLIENTS = 8
```
set to
```
	cFILE_PIPE_REJECT_REMOTE_CLIENTS = 0
```

### Reverse sock
client sock (init pipe connection:
```
main.exe -connect 127.0.0.1
```
server sock (listen pipe connection) :
```
main.exe -listen -socks 127.0.0.1:8080
```


### Simple sock
client sock (listen pipe connection) sock:
```
main.exe -listen 127.0.0.1
```
server sock (init pipe connection) :
```
main.exe  -connect 127.0.0.1 -socks 127.0.0.1:8080
```
