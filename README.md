[![pt-br](https://img.shields.io/badge/language-pt--br-green.svg)](https://github.com/kauemurakami/go-anonymous-functions/blob/main/README.pt-br.md)
[![en](https://img.shields.io/badge/language-en-orange.svg)](https://github.com/kauemurakami/go-anonymous-functions/blob/main/README.md)

## Anonymous functions
Basically it's a function that doesn't have a name, you invoke it in a different way.

### Starting project
Create a ```go-anonymous-functions``` directory with a ```main.go``` file with the default contents:  
```go
package main

func main() {

}
```
In our main we will call a function, but we will not declare it *explicitly*, see the example:  
```go
...
func main() {
	func() { // Will complain that it is not being used at this time.
		fmt.Println("Hello world")
	}
}
```
How do we execute it when it doesn't have a name? simple, we add parentheses ```()``` to the end of it:  
```go
...
func main() {
	func() { 
		fmt.Println("Hello world")
	}()
}
```
This way, what happens is the following, we declare the function so far:    
```go
func() { // Will complain that it is not being used at this time.
		fmt.Println("Hello world")
	}
```
Adding ```()``` to the end, we tell it to execute:  
```go
func() { 
		fmt.Println("Hello world") // output Hello world
	}()
```  
*Anonymous functions with parameters*  
```go
....
	text := "Text"
	func(text string) {
		fmt.Println(text) // output Text
	}(text)
}
```
*Anonymous functions with return*  
```go
	var result string = func(text string) string {
		return text
	}("My text to return")
	fmt.Println(result)
```
Here you just need to assign the result of the function to a variable to be able to use it.

