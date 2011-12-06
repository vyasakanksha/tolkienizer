package lexer

import (
   "testing"
   "os"
)

func TestThatThang(t *testing.T) {
<<<<<<< HEAD
    delim := make([]int, 4)
    delim[0] = ' '
    delim[1] = ','
    delim[2] = '.'
    delim[3] = 'þ'
    used := '.'
    f, _ := os.Open("samplefile")
    defer f.Close()
    dr := NewDelimReader(f, delim, used)

    atestslice := make([]byte, 4)
    num, err := dr.Read(atestslice)
    if (num != 4 || err != nil) {
        t.Fail()
    }
    if (string(atestslice) != "Mini") {
        t.Fail()
    }
    num, err = dr.Read(atestslice)
    if (num != 4 || err != nil || string(atestslice) != "m.lo") {
        t.Fail()
    }
    num, err = dr.Read(atestslice)
    if (num != 4 || err != nil || string(atestslice) != "ಠv") {
        t.Fail()
    }
    num, err = dr.Read(atestslice)
    if (num != 4 || err != nil) {
        t.Fail()
    }
    if (atestslice[0] != 'e' || atestslice[1] != 'l' || atestslice[2] != 224 || atestslice[3] != 178) {
        t.Fail()
    }
    num, err = dr.Read(atestslice)
    if (num != 4 || err != nil) {
        t.Fail()
    }
    if (atestslice[0] != 160 || string(atestslice[1:]) != "it.") {
        t.Fail()
    }
}
