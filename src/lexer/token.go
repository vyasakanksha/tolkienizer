package lexer

// A token type. Represents one "unit of input", could be anything that is part
// of a language, right now we support strings ( i.e. words ) and UTF-8 runes.
type Token interface {

   // Compare two tokens for equality. This function necessarily uses runtime
   // reflection to make sure the tokens are the same type.
   // return: 'true' if tokens are equal, 'false' otherwise.
   IsEqual( t Token ) bool
}

// The two current implementations of Token are for runes and string. They are
// declared here.
type runeToken int
type stringToken string

// Functions for runeToken 

// Creates new rune-based token type
// return: Token with underlying type '*runeToken'
func NewRuneToken( n int ) Token {
   i := runeToken( n )
   return &i
}

// Implement 'IsEqual()' to make runeToken a 'lexer.Token'
func ( this * runeToken ) IsEqual( t Token ) bool {

   // Check if the token we were given is a '*runeToken'
   if i, ok := t.(*runeToken); ok {
      // If so, compare it with 'this'
      return *this == *i
   } else {
      // Otherwise, they can't possibly be equal, return 'false'
      return false
   }

   panic( "Should never get here!" )
}

// Functions for stringToken 

// Creates new string-based token type
// return: Token with underlying type '*stringToken'
func NewStringToken( s string ) * stringToken {
   i := stringToken( s )
   return &i
}

// Implement 'IsEqual()' to make stringToken a 'lexer.Token'
func ( this * stringToken ) IsEqual( t Token ) bool {

   // Check if the token we were given is a '*stringToken'
   if i, ok := t.(*stringToken); ok {
      // If so, compare it with 'this'
      return *this == *i
   } else {
      // Otherwise, they can't possibly be equal, return 'false'
      return false
   }

   panic( "Should never get here!" )
}
