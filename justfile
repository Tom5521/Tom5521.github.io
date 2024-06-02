build-wasm:
  GOOS=js GOARCH=wasm go build -v -o my-page.wasm ./cmd/random-zazu-img/main.go
