//go:build js

package main

import (
	"math/rand"
	"syscall/js"
)

func main() {
	var images = []string{
		"./images/zazu/zazu.jpg",
		"./images/zazu/born-to-dilly-dally.jpg",
		"./images/zazu/theres-motion.jpg",
		"./images/zazu/zazu-in-the-war.jpg",
		"./images/zazu/zazu-in-watermelon.jpg",
	}

	randomID := rand.Intn(len(images))

	document := js.Global().Get("document")
	img := document.Call("getElementById", "zazu-img")
	img.Set("src", js.ValueOf(images[randomID]))
}
