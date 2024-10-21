package main


import (

	"fmt"
	"strconv"
	"unicode"
)


const (

	TOKEN_INT = "INT"
	TOKEN_PLUS = "+"
	TOKEN_MINUS   = "-"
	TOKEN_MUL     = "*"
	TOKEN_DIV     = "/"
	TOKEN_LPAREN  = "("
	TOKEN_RPAREN  = ")"
	TOKEN_EOF     = "EOF"
)


type Token struct{
	Type string 
	value string 
}

type Lexer struct{
	text  string
	pos   int 
	currentChar rune 
}


func NewLexer(text string) *Lexer{
	l := &Lexer{text:text, pos:0}
	l.currentChar = rune(text[l.pos])
	return l
}



func (l *Lexer) advance(){
	l.pos++
	if l.pos >= len(l.text){
		l.currentChar = 0 //End of input

	}else{
		l.currentChar = rune(l.text[l.pos])
	}
}



// skip whitesapce



func (l*Lexer) skipWhiteSpace(){
	for l.currentChar != 0  && unicode.IsSpace(l.currentChar){
		l.advance()
	}
}



//parse Integer Numbers

func (l *Lexer) integer string{
	result := ""
	for l.currentChar != 0 && unicode.IsDigit(l.currentChar){
		result += string(l.currentChar)
		l.advance()
	}
	return result
}