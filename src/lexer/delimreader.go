package lexer

import (
    "os"
    "io"
    "bufio"
    "utf8"
    "fmt"
)

type DelimReader struct {
    reader *(bufio.Reader)
    delimiters []int
    used_delimiter int
    remainder []byte
}

func NewDelimReader(read io.Reader, delim []int, used int) io.Reader {
    return &DelimReader{bufio.NewReader(read), delim, used, nil}
}

func (r *DelimReader) Read(p []byte) (n int, err os.Error) {
    bytes_written := 0
    for (bytes_written < len(p)) {
        fmt.Printf("Bites written: %d\n", bytes_written)
        rune, size, err := r.reader.ReadRune()
        if (err != nil) {
            return bytes_written, err
        }
        fmt.Printf("Rune read: %s\n", string(rune))
        for _, value := range r.delimiters {
            if (value == rune) {
                fmt.Printf("It's a delimiter\n")
                rune = r.used_delimiter
                size = utf8.RuneLen(rune)
            }
        }
        fmt.Printf("Rune size: %d\n", size)
        if (bytes_written + size > len(p)) {
            // we need to split the rune and hold on to the remainder
            writable := len(p) - bytes_written
            fmt.Printf("Splitting rune: %s, writable: %d\n", string(rune), writable)
            target := make([]byte, size)
            _ = utf8.EncodeRune(target, rune)
            for i := 0; i < writable; i++ {
                p[bytes_written] = target[i]
                fmt.Printf("%s %s\n", string(target[i]), string(p[bytes_written]))
                bytes_written++
            }
            fmt.Printf("%d\n", bytes_written)
            r.remainder = target[writable + 1:]
            fmt.Printf("%s %s\n", string(p), string(r.remainder))
        } else {
            fmt.Printf("Remaining bytes: %d\n", len(p) - bytes_written)
            target := p[bytes_written:bytes_written + size]
            _ = utf8.EncodeRune(target, rune)
            bytes_written += size
        }
    }
    return bytes_written, nil
}
