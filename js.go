// +build js,wasm

package main

import (
	"syscall/js"

	"github.com/greymd/ojichat/generator"
)

func main() {
	config := generator.Config{TargetName: "", EmojiNum: 4, PunctiuationLevel: 0}
	text, _ := generator.Start(config)
	js.Global().Get("document").Call("getElementById", "ojichatbox").Set("innerText", text)
}
