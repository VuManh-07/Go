package main

/*
#cgo windows CFLAGS: -DCGO_OS_WINDOWS=1
#cgo darwin CFLAGS: -DCGO_OS_DARWIN=1
#cgo linux CFLAGS: -DCGO_OS_LINUX=1

#if defined(CGO_OS_WINDOWS)
    const char* os = "windows";
#elif defined(CGO_OS_WINDOWS)
    const char* os = "windows";
#elif defined(CGO_OS_WINDOWS)
    const char* os = "windows";
#else
#   error(unknown os)
#endif
*/
import "C"

func main() {
	println(C.GoString(C.os))
}
