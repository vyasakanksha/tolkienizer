package pfa

import (
   "lexer"
   "sync"
)

// This package implements our "probabilistic finite automata", basically the
// same as Markov Processes. This name is kind of goofy, but I like it.

type PFA interface {
   Learn()
   GenerateSequence() string
}

type fnNetwork struct {
   input  lexer.NReader
   stream chan string
   mutex  *sync.RWMutex
}

func NewNetworkPFA(r lexer.NReader) PFA {
   ch := make(chan string, 0)
   return &fnNetwork{input: r, stream: ch, mutex: new(sync.RWMutex)}
}

func (this *fnNetwork) Learn() {

}

func (this *fnNetwork) GenerateSequence() string {
   return ""
}
