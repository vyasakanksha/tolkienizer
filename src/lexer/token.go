package lexer

//import "fmt"

type Token interface {
   Cmp( t Token ) bool
}

type intToken int
type stringToken string

// Functions for intToken 
func NewIntToken( n int ) * intToken {
   i := intToken( n )
   return &i
}

func ( this * intToken ) Cmp( t Token ) bool {
   if i, ok := t.(*intToken); ok {
      return *this == *i
   } else {
      return false
   }

   panic( "Should never get here!" )
}

// Functions for stringToken 
func ( this * stringToken ) Cmp( t Token ) bool {
   if i, ok := t.(*stringToken); ok {
      return *this == *i
   } else {
      return false
   }

   panic( "Should never get here!" )
}

func NewToken( s string ) * stringToken {
   i := stringToken( s )
   return &i
}

/*func main() {
   i := NewIntToken( 5 )
   j := NewIntToken( 5 )
   k := NewIntToken( 6 )

   fmt.Println( i.Cmp( j ) )
   fmt.Println( i.Cmp( k ) )
}*/
