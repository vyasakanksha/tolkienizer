package lexer

import (
   "bufio"
)

type NReader interface {

   // This function returns the last n tokens in the tokenSet
   Prefix() []Token

   // This function returns the most recent token in the tokenSet
   Next() Token

   // This function reads a new token and adds it to tokenSet
   Advance()
}

type NReaderBase struct {
   n        int
   next     Token
   tokenSet []Token
   r        bufio.Reader
   negOne   Token
}

// This function looks at the righmost value in the tokenset and returns a slice
// containing the last n tokens
func (this NReaderBase) Prefix() []Token {
   j := 0
   length := len(this.tokenSet)
   prefixSet := make([]Token, this.n)

   if length < this.n {
      for j := 0; j < (this.n - length); j++ {
         prefixSet[j] = this.negOne
      }
   }

   for i := (length - this.n); i < length; i++ {
      prefixSet[j] = this.tokenSet[i]
      j++
   }
   return prefixSet
}

//This function reads a new value, converts it into a token and adds it to
//tokenSet
/*
func ( this NReaderBase ) Advance() {
  temp, _, err := this.r.ReadRune()
}
*/

// This function returns the last token from tokenSet
func (this NReaderBase) Next() Token {
   return this.tokenSet[len(this.tokenSet)-1]
}

//type NReaderSpace struct {
//NReaderBase
//}
