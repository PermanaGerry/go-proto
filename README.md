# go-proto
Project ini merupakan example implementasi dari protocol buffers dan grpc dengan menggunakan bahasa go. Untuk menjalankan protobuff harus terinstall applikasinya terlebih dahulu.

## Installation on windows
Bisa menggunakan link ini [geeksforgeeks] (https://www.geeksforgeeks.org/how-to-install-protocol-buffers-on-windows/)

## Installation on linux

```
$ apt install -y protobuf-compiler
$ protoc --version
```
## Installation on macos

```
$ brew install protobuf
$ protoc --version  # Ensure compiler version is 3+
```

## Usage

```
# run protoc
protoc *.proto --go_out=.

# run application go
go run cmd/main.go
```