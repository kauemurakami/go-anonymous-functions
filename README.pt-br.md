[![pt-br](https://img.shields.io/badge/language-pt--br-green.svg)](https://github.com/kauemurakami/go-anonymous-functions/blob/main/README.pt-br.md)
[![en](https://img.shields.io/badge/language-en-orange.svg)](https://github.com/kauemurakami/go-anonymous-functions/blob/main/README.md)

## Funções anônimas
Basicamente é uma função que não tem nome, você invoca ela de um jeito diferente.  

### Iniciando projeto
Crie um diretório ```go-anonymous-functions``` com um arquivo ```main.go``` com o conteúdo padrão:  
```go
package main

func main() {

}
```
Em nossa main vamos chamar uma função, mas não vamos declarar ela *explicitamente*, veja no exemplo:  
```go
...
func main() {
	func() { // Irá reclamar que ela não está sendo usada neste momento.
		fmt.Println("Olá mundo")
	}
}
```
Como fazemos para executá-la sendo que ela não possui nome? simples, adicionamos parênteses ```()``` ao inal dela:  
```go
...
func main() {
	func() { 
		fmt.Println("Hello world")
	}()
}
```
Dessa forma o que acontece é o seguinte, declaramos a função até aqui:  
```go
func() { // Irá reclamar que ela não está sendo usada neste momento.
		fmt.Println("Hello world")
	}
```
Adicionando o ```()``` ao final, dizemos para executar ela:  
```go
func() { 
		fmt.Println("Hello world") // output Hello world
	}()
```  
*Funções anônimas com parâmetros*  
```go
....
	text := "Text"
	func(text string) {
		fmt.Println(text) // output Text
	}(text)
}
```
*Funções anônimas com retorno*  
```go
	var result string = func(text string) string {
		return text
	}("My text to return")
	fmt.Println(result)
```
Aqui basta atribuir o resultado da função em uma variável para poder usa-lá.  

