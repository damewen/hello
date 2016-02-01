package main

import (
	"github.com/spakin/awk"
	"os"
	"fmt"
)

func main() {
	fmt.Printf("affichage de la colonne 2 \n")
	arg := os.Args[1]
	fi, err := os.Open(arg)
	if err != nil {
        panic(err)
    }
	s := awk.NewScript()
	//s.AppendStmt(awk.Auto("yoann"), nil)
	s.SetFS(";")
	s.AppendStmt(nil, func(s *awk.Script) { s.Println(s.F(2)) })
	s.Run(fi)
}