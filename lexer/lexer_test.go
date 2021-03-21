package lexer

import (
	"testing"

	"github.com/ECAllen/mymarkdown/token"
)

func TestNextToken(t *testing.T){
	input := ` card:
x: y
`

	tests := [ ]struct {
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.CARD, "card"},
		{token.COLON, ":"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q,", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
