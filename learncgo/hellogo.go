package main

// #include "hellogo.h"
import "C"

import "fmt"

// export PrintHello
func PrintHello() {
	fmt.Println(C.GoString("Hello"))
}
