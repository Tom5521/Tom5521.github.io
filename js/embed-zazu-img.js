const go = new Go();
WebAssembly.instantiateStreaming(
  fetch("./wasm/random-zazu-img.wasm"),
  go.importObject,
).then((result) => {
  go.run(result.instance);
});
