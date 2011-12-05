package lexer

import (
   "testing"
)

func TestStringIsEqual(t *testing.T) {
   i := NewStringToken("thanks for all")
   j := NewStringToken("thanks for all")
   k := NewStringToken("fish")

   if !i.IsEqual(j) {
      t.Errorf("String %s != %s", string(*i), string(*j))
   } else if i.IsEqual(k) {
      t.Errorf("String %s == %s", string(*i), string(*k))
   }
}
