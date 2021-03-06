package lexer

import (
   "os"
   "io"
   "bufio"
   "utf8"
)

type DelimReader struct {
   reader         *(bufio.Reader)
   delimiters     []int
   used_delimiter int
   remainder      []byte
}

func NewDelimReader(read io.Reader, delim []int, used int) io.Reader {
   return &DelimReader{bufio.NewReader(read), delim, used, nil}
}

func (r *DelimReader) Read(p []byte) (n int, err os.Error) {
   bytes_written := 0

   if r.remainder != nil {
      for i := 0; i < len(r.remainder); i++ {
         p[i] = r.remainder[i]
         bytes_written++
      }
      r.remainder = nil
   }

   for bytes_written < len(p) {
      rune, size, err := r.reader.ReadRune()
      if err != nil {
         return bytes_written, err
      }
      for _, value := range r.delimiters {
         if value == rune {
            rune = r.used_delimiter
            size = utf8.RuneLen(rune)
         }
      }
      if bytes_written+size > len(p) {
         // we need to split the rune and hold on to the remainder
         writable := len(p) - bytes_written
         target := make([]byte, size)
         _ = utf8.EncodeRune(target, rune)
         for i := 0; i < writable; i++ {
            p[bytes_written] = target[i]
            bytes_written++
         }
         r.remainder = target[writable:]
      } else {
         target := p[bytes_written : bytes_written+size]
         _ = utf8.EncodeRune(target, rune)
         bytes_written += size
      }
   }
   return bytes_written, nil
}
