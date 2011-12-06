package pfa

import "testing"

var (
   stringSets = [][]string{
      {"a", "b", "c"},
      {"a", "b", "d"},
      {"a", "c", "e"},
      {"f", "z", "h"}}
)

func TestInsertAndLookup(t *testing.T) {
   tt := newTransitionTable()

   dummy := make(chan string)

   for i := range stringSets {
      if err := tt.insert(stringSets[i], dummy); err != nil {
         t.Errorf("Was unable to insert with key '%v'.", stringSets[i])
      }
   }

   for i := range stringSets {
      if ch, err := tt.lookUp(stringSets[i]); err != nil || ch != dummy {
         t.Errorf("Lookup failed on '%v'.", stringSets[i])
      }
   }
}
