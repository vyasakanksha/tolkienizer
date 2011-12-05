package lexer

import (
   "bufio"
)

type NReader interface {

   // This function returns the last n strings in the tokenSet
   Prefix() []string

   // This function returns the most recent string in the tokenSet
   Next() string

   // This function reads a new string and adds it to tokenSet
   Advance()
}

// A base struct off which each reader extends
type nReaderBase struct {
   // The number of strings each prefix contains
   n  int

   // The leftmost string in the token set
   next string

   // A slice of strings in the order they appear in the learning data
   stringSet []string

   // Specifies the format in which the learning data is stored
   r  bufio.Reader
}

type nReaderSpace struct {
   nReaderBase
}

type nReaderNoSpace struct {
   nReaderBase
}

type nReaderWords struct {
   nReaderBase
}

// This function looks at the righmost value in the tokenset and returns a slice
// containing the last n tokens
func (this NReaderBase) Prefix() []string {
   j := 0
   length := len(this.stringSet)
   prefixSet := make([]string, this.n)

   if length < this.n {
      for j := 0; j < (this.n - length); j++ {
         prefixSet[j] = "-1"
      }
   }

   for i := (length - this.n); i < length; i++ {
      prefixSet[j] = this.stringSet[i]
      j++
   }
   return prefixSet
}

// This function returns the last string from stringSet
func (this NReaderBase) Next() string {
   return this.stringSet[len(this.stringSet)-1]
}

//This function reads a new value, converts it into a string and adds it to
//stringSet
func (this NReaderSpace) Advance() {
   temp, _, err := this.r.ReadRune()
}
