package main

/*
void PrintHello(void);
*/
import "C"

import "fmt"

func main() {
	C.PrintHello()
}

//export PrintHello
func PrintHello() {
	fmt.Println("Hello")
}
