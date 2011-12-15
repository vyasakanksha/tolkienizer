package lexer

import (
   "testing"
   "strings"
)

func TestThatThang(t *testing.T) {
   delim := []int{' ', ',', '.', 'þ'}

   r := strings.NewReader("Minim loಠvelಠit,")

   dr := NewDelimReader(r, delim, '.')

   atestslice := make([]byte, 4)
   num, err := dr.Read(atestslice)
   if num != 4 || err != nil {
      t.Fail()
   }
   if string(atestslice) != "Mini" {
      t.Fail()
   }
   num, err = dr.Read(atestslice)
   if num != 4 || err != nil || string(atestslice) != "m.lo" {
      t.Fail()
   }
   num, err = dr.Read(atestslice)
   if num != 4 || err != nil || string(atestslice) != "ಠv" {
      t.Fail()
   }
   num, err = dr.Read(atestslice)
   if num != 4 || err != nil {
      t.Fail()
   }
   if atestslice[0] != 'e' || atestslice[1] != 'l' || atestslice[2] != 224 || atestslice[3] != 178 {
      t.Fail()
   }
   num, err = dr.Read(atestslice)
   if num != 4 || err != nil {
      t.Fail()
   }
   if atestslice[0] != 160 || string(atestslice[1:]) != "it." {
      t.Fail()
   }
}
