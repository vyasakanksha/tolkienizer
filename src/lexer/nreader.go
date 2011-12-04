package lexer

import (
   "bufio"
)

type NReader interface {

   // This function returns the last n tokens in the tokenSet
   Prefix() []string

   // This function returns the most recent token in the tokenSet
   Next() string

   // This function reads a new token and adds it to tokenSet
   Advance()
}

type NReaderBase struct {
   n int
   next string
   stringSet []string
   r bufio.Reader
}


// This function looks at the righmost value in the tokenset and returns a slice
// containing the last n tokens
func ( this NReaderBase ) Prefix() []string {
   j := 0
   length := len( this.stringSet )
   prefixSet := make( []string, this.n )

   if length < this.n {
      for j := 0; j < (this.n - length); j++ {
         prefixSet[j] = "-1"
      }
   }

   for i := ( length - this.n ); i < length; i++ {
      prefixSet[j] = this.stringSet[i]
      j++
   }
   return prefixSet
}

//This function reads a new value, converts it into a string and adds it to
//stringSet
/*
func ( this NReaderBase ) Advance() {
  temp, _, err := this.r.ReadRune()
}
*/

// This function returns the last string from stringSet
func ( this NReaderBase ) Next() string {
   return this.stringSet[ len( this.stringSet ) - 1 ]
}

