package main

import (
	"errors"
	"fmt"
)

func main() {
	//var a string
	//fmt.Print(a)
	//fmt.Print(a == "")
	er := errors.New("wew")
	fmt.Print(er)
}
