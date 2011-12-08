package lexer

import (
   "testing"
   "strings"
)

func nReadersTest(t *testing.T) {
   // Test 1: nReaderLetter
   r := strings.NewReader("h ah")
   temp := NewNReaderLetter( 3, r )

   if temp.Advance() != false {

      test := []string{"","","h"}

      for i:= 0; i < len(test); i++  {
         if temp.Prefix()[i] != test[i] {
         t.Errorf("Expected \"h\" != %s", temp.Prefix() )
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
