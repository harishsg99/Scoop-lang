package lexer

import {
    "testing"

    "scooplang/token"
}

func TextNextToken(t *testing.T){
    input : = '=+-^/(){},;'

    tests :=[]struct {
        exceptedtype token.TokenType
        exceptedLiteral string
    }{
{token.ASSIGN, "="},
{token.PLUS, "+"},
{token.LPAREN, "("},
{token.RPAREN, ")"},
{token.LBRACE, "{"},
{token.RBRACE, "}"},
{token.COMMA, ","},
{token.SEMICOLON, ";"},
{token.MINUS, "-"},
{token.DIVIDE, "/"},
{token.POWER, "^"},
{token.MULTIPLY, "*"},
{token.EOF, ""},
}
l := New(input)


}
for i, tt := range tests {
tok := l.NextToken()

if tok.Type != tt.expectedType {
t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",i, tt.expectedType, tok.Type)
}
if tok.Literal != tt.expectedLiteral {
t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
i, tt.expectedLiteral, tok.Literal)
}
}
}
