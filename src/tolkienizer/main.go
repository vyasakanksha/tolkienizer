package main

import (
   "fmt"
   "lexer"
   "pfa"
   "os"
)

func main() {
   fmt.Printf("Tolkienizer v0.0\n")
   nr := lexer.NewNReaderLetterDelim(3, os.Stdin, []int{' ', '.', ',', ';',
      '\n', '\'', '"', '`', ':', '-', '\t'})

   tpfa := pfa.NewTablePFA(nr)

   tpfa.Learn()

   for i := 0; i < 20; i++ {
      fmt.Println(tpfa.GenerateString())
   }
}
