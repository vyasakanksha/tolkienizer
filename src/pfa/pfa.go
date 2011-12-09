package pfa

import (
   "strings"
   "lexer"
   "rand"
   "time"
)

type PFA interface {
   Learn()
   GenerateString() string
}

type tablePfa struct {
   input lexer.NReader
   table map[string]*tablePfaState
   random  *rand.Rand
}

type tablePfaState struct {
   prefix     []string
   edges      map[string]int
   totalCount int
}

func NewTablePFA(r lexer.NReader) PFA {
   return &tablePfa{
      input: r,
      table: make(map[string]*tablePfaState),
      random: rand.New( rand.NewSource( time.Nanoseconds() ) )}
}

func (me *tablePfa) Learn() {
   for keepGoing := me.input.Advance(); keepGoing; keepGoing = me.input.Advance() {

      key := strings.Join(me.input.Prefix(), ":")

      if _, ok := me.table[key]; !ok {

         me.table[key] = &tablePfaState{
            prefix: me.input.Prefix(),
            edges: make(map[string]int),
            totalCount: 0}

      }

      me.table[key].edges[me.input.Next()]++
      me.table[key].totalCount++
   }
}

func (me *tablePfa) GenerateString() string {
   // Generate the start key, ":::" if N = 3
   key := strings.Join(make([]string, me.input.GetN()), ":")
   ret := ""
   for state := me.table[key]; state != nil ; state = me.table[key] {
      rNum := me.random.Intn(state.totalCount)
      for k, v := range state.edges {
         rNum -= v
         if rNum < 1 {
            if k == "" {
               return ret
            } else {
               ret += k
               key = strings.Join(append(state.prefix[1:], k), ":")
               break
            }
         }
      }
   }
   return ""
}
