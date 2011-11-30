package lexer

import(
   "os"
   "fmt"
   )

type Token struct {
   iSet *int
   sSet *string
}

func Cmp( a Token, b Token ) bool {
   switch {
      case a.iSet == nil && b.iSet == nil:
            if a.sSet == nil || b.sSet == nil {
               fmt.Fprintf( os.Stderr, "Invalid Input" )
            } else {
               if &a == &b {
                  return true
               } else { return false }
            }
      case a.sSet == nil && b.sSet == nil:
            if a.iSet == nil || b.iSet == nil {
               fmt.Fprintf( os.Stderr, "Invalid Input" )
            } else {
               if &a == &b {
                  return true
               } else { return false }
            }
   }

  fmt.Fprintf( os.Stderr, "Inconsistent Types" )
  return false
}
