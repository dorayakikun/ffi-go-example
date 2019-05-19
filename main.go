package main

import (
	"C"
	"fmt"
	"os"
)
//export print
func print(cstr *C.char) {
	text := C.GoString(cstr)
	fmt.Fprint(os.Stdout, text)
}

func main() {}