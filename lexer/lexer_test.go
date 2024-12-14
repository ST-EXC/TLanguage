package lexer

import (
	"TLanguage/token"
	"fmt"
	"testing"
)

func inputProcedure(c int) string {
	var input string
	switch c {
	case 0:
		input = `=+(){},;"foobar"`
	case 1:
		input = `let five = 5;
		let ten = 10;
		let add = fn(x,y){
		   x + y;
		};
		let result = add(five,ten)`
	case 2:
		input = `let five = 5;
		let ten = 10;
		let add = fn(x,y){
			let result = x+y;
			return result;
		}
		let result = add(five,ten);
		!=/*5
		5<10>5
		if (5<10){
			return true;
		}else{
			return false;
		}
		10==10
		10!=9
		`
	}
	return input
}

func TestNextToken(t *testing.T) {
	input := inputProcedure(0)
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.STRING, "foobar"},
		{token.EOF, ""},
	}
	l := NewLexer(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. Expected %q, got %q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. Expected %q, got %q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken2(t *testing.T) {
	fmt.Println("TEST START")
	t.Logf("start")
	input := inputProcedure(2)
	l := NewLexer(input)
	for {
		tok := l.NextToken()
		if tok.Type == token.EOF {
			break
		}
		t.Logf("tokentype %q, literal %q", tok.Type, tok.Literal)
	}
}
