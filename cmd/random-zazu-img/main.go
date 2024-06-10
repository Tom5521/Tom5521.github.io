//go:build wasm

package main

import (
	"embed"
	"log"
	"math/rand"
	"syscall/js"
)

//go:embed zazu
var imagesFS embed.FS

func main() {
	dirs, err := imagesFS.ReadDir("zazu")
	if err != nil {
		log.Panic(err)
	}

	var images []string

	for _, dir := range dirs {
		images = append(images, "./images/zazu/"+dir.Name())
	}

	randomID := rand.Intn(len(images))

	document := js.Global().Get("document")
	img := document.Call("getElementById", "zazu-img")
	img.Set("src", js.ValueOf(images[randomID]))
}
