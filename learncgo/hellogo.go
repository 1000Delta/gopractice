package main

/*
#include "hello.h"
*/
import "C"
import "fmt"

//export PrintHelloGo
func PrintHelloGo() {
	fmt.Println(C.GoString(C.CString("Hello Go")))
}
