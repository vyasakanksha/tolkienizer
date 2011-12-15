package lexer

import (
   "testing"
   "strings"
)

func TestNReaders(t *testing.T) {
   // Test 1: nReaderLetter
   r := strings.NewReader("h ah")
   temp := NewNReaderLetter(3, r)

   if temp.Advance() != false {

      test := []string{"", "", ""}

      for i, rune := range temp.Prefix() {
         if rune != test[i] {
            t.Errorf("Expecting \"%#v\", got \"%#v\"", test, temp.Prefix())
         }
      }
   }
}

/* delim := {' '}
   newNReaderLettersDelim(3, r, ' ')
   Advance()

   newNReaderWords(3, r, ' ')
   Advance()
*/
