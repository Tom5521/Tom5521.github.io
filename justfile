run-server:
    go run -v ./cmd/server/main.go

build-wasm:
    rm -rf wasm
    GOOS=js GOARCH=wasm go build -v -o ./wasm/random-zazu-img.wasm ./cmd/random-zazu-img/main.go
