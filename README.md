# ファイル交換用

受信側はサーバーを起動する必要があります。

## サーバーの使用方法
./file_exchange

## clientの使用方法
./file_exchange 対象のip ファイル名

## compile方法

### Windows
**32bit**  
`GOOS=windows GOARCH=386 go build file_exchange.go`

**64bit**  
`GOOS=windows GOARCH=amd64 go build file_exchange.go`

### Mac
**32bit**  
`GOOS=darwin GOARCH=386 go build file_exchange.go`

**64bit**  
`GOOS=darwin GOARCH=amd64 go build file_exchange.go`

### Linux
`GOOS=linux GOARCH=amd64 go build file_exchange.go`

