build-wasm:
  GOOS=js GOARCH=wasm go build -v -o ./wasm/random-zazu-img.wasm ./cmd/random-zazu-img/main.go
build-js:
  fyne package -os js
