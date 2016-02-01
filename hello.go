package main

import (
	"github.com/spakin/awk"
	"os"
	"fmt"
)

func main() {
	fmt.Printf("search for \"yoann\" \n")
	arg := os.Args[1]
	fi, err := os.Open(arg)
	if err != nil {
        panic(err)
    }
	s := awk.NewScript()
	s.AppendStmt(awk.Auto("yoann"), nil)
	s.Run(fi)
}