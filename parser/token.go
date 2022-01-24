package parser

// A token in a Bencoding file.
//
// value is set to the token's string/integer for the stringToken/intToken types, or to the empty string otherwise.
type token struct {
	tokenType tokenType
	value     []byte
}

// The type of the token.
type tokenType uint8

// The types of token.
const (
	stringToken tokenType = iota
	intToken
	listStartToken
	dictStartToken
	endToken
)
