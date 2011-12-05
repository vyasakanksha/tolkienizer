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
   Advance() bool
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

// Type that treats every letter (including punctuation and space) as a token
type nReaderLetter struct {
   nReaderBase
}

// Type that treats every letter as a token, and treats every word as
// a different problem. I.e. everytime it encounters a space or punctuation, 
// it stops reading.
type nReaderLetterDelim struct {
   nReaderBase
}

// Type that treats each word as a token
type nReaderWords struct {
   nReaderBase
}

func (this nReaderLetterDelim) Advance() bool {
   temp, _, err := this.r.ReadRune()
   if err != nil || temp == '.' {
      return false
   } else {
      this.stringSet[ len(this.stringSet) ] = string(temp)
      return true
   }

   panic( "We should never get here")
}


func (this nReaderLetter) Advance() bool {
   temp, _, err := this.r.ReadRune()
   if err != nil {
      return false
   } else {
      this.stringSet[len(this.stringSet)] = string(temp)
      return true
   }

   panic( "We should never get here")
}

func (this nReaderWords) Advance() bool {
   temp, err := this.r.ReadString('.')
   if err != nil {
      return false
   } else {
      this.stringSet[len(this.stringSet)] = temp
      return true
   }

   panic( "We should never get here")
}

// Return: Last n values of the stringSet (excuding the current value)
func (this nReaderBase) Prefix() []string {
   j := 0
   length := len(this.stringSet)
   prefixSet := make([]string, this.n)

   // If n is larger than the stringSet, initialize prefixSet[:n-length] to
   // the empty string
   if length < this.n {
      for j := 0; j < (this.n - length); j++ {
         prefixSet[j] = ""
      }
   }

   // Copy the last n (excluding length) tokens to prefix
   for i := (length - this.n); i < length; i++ {
      prefixSet[j] = this.stringSet[i]
      j++
   }
   return prefixSet
}

// This function returns the current string from stringSet
func (this nReaderBase) Next() string {
   return this.stringSet[len(this.stringSet)-1]
}

