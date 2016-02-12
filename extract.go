package main

import (
	"github.com/spakin/awk"
	"os"
	"fmt"
	"strconv"
)

func main() {

	
	arg := os.Args[1]
	col_str := os.Args[2]
	fi, err := os.Open(arg)
	col_int, err := strconv.Atoi(col_str)
	if err != nil {
        panic(err)
    }
	fmt.Printf("affichage de la colonne %d \n", col_int)
	s := awk.NewScript()
	s.SetFS(";")
	s.AppendStmt(nil, func(s *awk.Script) { s.Println(s.F(col_int)) })
	s.Run(fi)
}