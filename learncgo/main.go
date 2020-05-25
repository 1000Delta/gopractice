package main

/*
#include "hello.h"
void PrintHelloGo(void);
*/
import "C"

func main() {
	C.PrintHelloC()
	C.PrintHelloGo()
}
