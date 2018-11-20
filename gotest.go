package main

import (
	"fmt"
	"github.com/qiniu/x/errors.v7"
)

func main() {
	//var a string
	//fmt.Print(a)
	//fmt.Print(a == "")
	er := errors.New("wew")
	fmt.Print(er)
}
