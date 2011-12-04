package main

/* import "fmt" */

// An interface for Token that can be of any type as long it implements 
// Cmp (compare). 
type Token interface {
   Cmp( t Token ) bool
}

// Two current implementions of Token are for ints and string. They are declared
// here.
type intToken int
type stringToken string

// Functions for intToken 

// Default constructor - takes an integer as input, converts it into a intToken
//and returns a pointer to the intToken
func NewIntToken( n int ) * intToken {
   i := intToken( n )
   return &i
}

// This function takes a token as input and compares it to an intToken,
// returning true if they are equal. It first checks if the input token is 
// indeed an intToken using type reflection.
// The func returns false on incompatable types and inequality
func ( this * intToken ) Cmp( t Token ) bool {
   if i, ok := t.(*intToken); ok {
      return *this == *i
   } else {
      return false
   }

   panic( "Should never get here!" )
}

// Functions for stringToken 

// Default constructor - takes an String as input, converts it into a
// stringToken and returns a pointer to the stringToken.
func NewStringToken( s string ) * stringToken {
   i := stringToken( s )
   return &i
}

// This function takes a token as input and compares it to a stringToken,
// returning true if they are equal. It first checks if the input token is 
// indeed an stringToken using type reflection.
// The func returns false on incompatable types and inequality
func ( this * stringToken ) Cmp( t Token ) bool {
   if i, ok := t.(*stringToken); ok {
      return *this == *i
   } else {
      return false
   }

   panic( "Should never get here!" )
}

/*func main() {
   i := NewStringToken( "thanks for all" )
   j := NewStringToken( "thanks for all" )
   k := NewStringToken( "fish" )

   fmt.Println( i.Cmp( j ) )
   fmt.Println( i.Cmp( k ) )
}*/
