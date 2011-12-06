package lexer

import (
    "testing"
    "fmt"
)

func TestThatThang(t *testing.T) {
    delim := make([]int, 4)
    delim[0] = ' '
    delim[1] = ','
    delim[2] = '.'
    delim[3] = '-'
    used := '.'
    dr := NewDelimReader("shittysamplefile", delim, used)

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
    if (num != 4 || err != nil || string(atestslice) != "elà²") {
        fmt.Printf("%s\n", atestslice)
        t.Fail()
    }
}
