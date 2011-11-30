package lexer

type NReader interface {

   // returns the last n tokens for the current Token
   Prefix() []Token

   // returns the next Token for a slice of n tokens
   Next Token()

   // Reads a Token from the reader and appends it to the tokenSet
   Advance()
}

type NReaderBase {
   n int
   next Token
   tokenSet []Token
   r Reader
}

func ( NReaderBase ) Prefix() []token.Token {
   j := 0
   length := len( tokenSet )
   prefixSet := make( []Token, length - n )
   for i := ( length - n ); i < length; i++ {
      prefixSet[j] = tokenSet[i]
      j++
   }
}

type NReaderSpace struct {
   NReaderBase
}
