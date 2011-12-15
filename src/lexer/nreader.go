package lexer

import (
   "bufio"
   "io"
)

type NReader interface {

   // This function returns the last n strings in the tokenSet
   Prefix() []string

   // This function returns the most recent string in the tokenSet
   Next() string

   // This function reads a new string and adds it to tokenSet
   Advance() bool

   GetN() int
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
   r  *bufio.Reader
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

   // A list of characters that should as delimiters between words
   delim []int
}

// Type that treats each word as a token
type nReaderWords struct {
   nReaderBase

   // A list of characters that should as delimiters between words
   delim []int
}

// Initializes NReaderLetter with n and reader r 
func NewNReaderLetter(n int, r io.Reader) NReader {
   return &nReaderLetter{
      nReaderBase: nReaderBase{
         stringSet: make([]string, n),
         r:         bufio.NewReader(r),
         n:         n}}
}

// Initializes NReaderLetterDelim with n, reader r and a slice of delims that
// indicate the end of a word.
func NewNReaderLetterDelim(n int, r io.Reader, delim []int) NReader {
   return &nReaderLetterDelim{
      nReaderBase: nReaderBase{
         stringSet: make([]string, n),
         r:         bufio.NewReader(NewDelimReader(r, delim, '.')),
         n:         n},
      delim: delim}
}

// Initializes NReaderWords with n, reader r and a slice of delims that
// indicate the end of a word.
func NewNReaderWords(n int, r io.Reader, delim []int) NReader {
   return &nReaderWords{
      nReaderBase: nReaderBase{
         stringSet: make([]string, n),
         r:         bufio.NewReader(NewDelimReader(r, delim, '.')),
         n:         n},
      delim: delim}
}

// Return: Last n values of the stringSet (excuding the current value)
func (this *nReaderBase) Prefix() []string {
   return this.stringSet[len(this.stringSet)-(this.n+1) : len(this.stringSet)-1]
}

// This function returns the current string from stringSet
func (this *nReaderBase) Next() string {
   return this.stringSet[len(this.stringSet)-1]
}

func (this *nReaderBase) GetN() int {
   return this.n
}

// Implementation of the advance function that reads letters till it encounters
// a specified delimited. The function reads the next character from the stream
// and adds it to stringSet.
func (this *nReaderLetterDelim) Advance() bool {
   temp, _, err := this.r.ReadRune()

   // Since the filter converted all delims to ".", we only have to account for
   // it.
   if err != nil {
      return false
   } else if this.Next() == "" {
      for ; temp == '.' && err == nil; temp, _, err = this.r.ReadRune() {
      }
      this.stringSet = make([]string, this.n+1)
      this.stringSet[len(this.stringSet)-1] = string(temp)
   } else if temp == '.' {
      this.stringSet = append(this.stringSet, "")
   } else {
      this.stringSet = append(this.stringSet, string(temp))
   }

   return true
}

// Implementation of the advance function that keeps reading letters till 
// it encounters the end of the stream. The function reads the next 
// character from the stream and adds it to stringSet.
func (this *nReaderLetter) Advance() bool {
   temp, _, err := this.r.ReadRune()
   if err != nil {
      return false
   } else {
      this.stringSet = append(this.stringSet, string(temp))
      return true
   }

   panic("We should never get here")
}

// Implementation of the advance function that reads stings (words) till 
// it encounters a specified delimiter. The function reads the next 
// string from the stream and adds it to stringSet.
func (this *nReaderWords) Advance() bool {
   // Since the filter converted all delims to ".", we only have to account for
   // it.
   temp, err := this.r.ReadString('.')
   if err != nil {
      return false
   } else {
      this.stringSet = append(this.stringSet, temp)
      return true
   }

   panic("We should never get here")
}
