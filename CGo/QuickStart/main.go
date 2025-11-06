package main

/*
#include <stdlib.h>
#include "hello.h" // PHẢI CÓ dòng này
*/
import "C"

func main() {
	C.SayHello(C.CString("Hello World!"))
}

// go run . : de compile all file
