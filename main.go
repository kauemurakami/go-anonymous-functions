package main

import "fmt"

func main() {
	// função anônima
	func() {
		fmt.Println("Hello world")
	}()
	// função anônima com parâmetros
	text := "Text"
	func(text string) {
		fmt.Println(text)
	}(text)
	// função anônima com retorno
	var result string = func(text string) string {
		return text
	}("My text to return")
	fmt.Println(result)
}
