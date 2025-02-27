//go:build js && wasm

package main

import (
	"syscall/js"

	"wasm-calculator/calculator"
)

func main() {
	c := &calculator.Calculator{}

	// Определяем функции для JavaScript
	js.Global().Set("wasmAdd", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 2 {
			return 0
		}
		return c.Add(args[0].Float(), args[1].Float())
	}))

	js.Global().Set("wasmSubtract", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 2 {
			return 0
		}
		return c.Subtract(args[0].Float(), args[1].Float())
	}))

	js.Global().Set("wasmMultiply", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 2 {
			return 0
		}
		return c.Multiply(args[0].Float(), args[1].Float())
	}))

	js.Global().Set("wasmDivide", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 2 {
			return 0
		}
		return c.Divide(args[0].Float(), args[1].Float())
	}))

	// Предотвращаем завершение программы
	select {}
}
