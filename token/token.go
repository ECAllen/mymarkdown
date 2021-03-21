package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"
    CARD = "CARD"
	COLON = ":"
	IDENT = "IDENT"
	SEMICOLON = "\n"
	INT = "INT"
)

var keywords = map[string]TokenType{
	"card": CARD,
}

func LookupIdent(ident string) TokenType{
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
