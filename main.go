package main

import (
	"syscall/js"
)

// The phrase we want to type out.
const targetPhrase = "This is a message being typed by Go WebAssembly!"

// A variable to track our position in the phrase.
var currentIndex = 0

// This function will be our event handler for key presses.
func typeNextCharacter(this js.Value, args []js.Value) interface{} {
	// The first argument is the JavaScript event object.
	event := args[0]

	// This is crucial: it stops the browser from typing the actual
	// character that the user pressed.
	event.Call("preventDefault")

	// If we haven't finished typing the phrase yet...
	if currentIndex < len(targetPhrase) {
		// Find the text box element in the HTML.
		textBox := js.Global().Get("document").Call("getElementById", "emailBox")

		// Get the next character from our phrase.
		nextChar := string(targetPhrase[currentIndex])

		// Append our character to the text box's current value.
		currentValue := textBox.Get("value").String()
		textBox.Set("value", currentValue+nextChar)

		// Move to the next character for the next key press.
		currentIndex++
	}

	return nil
}

func main() {
	// Create a channel that will never receive a value.
	// This keeps the Go program alive so it can listen for events.
	done := make(chan struct{}, 0)

	// Get the HTML document object.
	document := js.Global().Get("document")

	// Find our text box element by its ID.
	emailBox := document.Call("getElementById", "emailBox")

	// Create a Go function that can be called by JavaScript.
	jsFunc := js.FuncOf(typeNextCharacter)

	// Attach the Go function as an event listener for the "keydown" event.
	emailBox.Call("addEventListener", "keydown", jsFunc)

	// Wait forever.
	<-done
}
