package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("This is a exercise of cross compilation. This program was compiled in", runtime.GOARCH, runtime.GOOS)
}

//Change compile OS
//GOOS=darwin GOARCH=amd64 go build compile.go
