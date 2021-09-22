package main

import (
	"fmt"
	"os"

	"github.com/MartIden/deep-throtle-parser/command"
)

func main() {
	_,err:= command.ProcessArgs(os.Args)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

}