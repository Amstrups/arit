package lexer

import (
	par "arit/parsing"
	"fmt"
	"strings"
	"testing"
)

type ASSERT_ERROR_CODE int

const (
	NO_ERROR = iota
	ELEMENT_INEQUALITY
	LENGTH_INEQAULITY
)

type AssertError struct {
	code ASSERT_ERROR_CODE
	msg  string
}

func toString[T par.Token](ts []T, fn func(T) string) []string {
	result := make([]string, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

func concat(xs []par.Token) string {
	ss := toString(xs, func(t par.Token) string { return t.S })

	return "[" + strings.Join(ss, ",") + "]"
}

func concat2(xs []par.TokenCode) string {
	ss := make([]string, len(xs))
	for i, x := range xs {
		ss[i] = fmt.Sprintf("%d", x)
	}

	return "[" + strings.Join(ss, ",") + "]"
}

func prepareMessage(ts []par.Token, tcs []par.TokenCode) string {

	return fmt.Sprintf("\n%10s %s\n%10s %s", "Found:", concat(ts), "Expected:", concat2(tcs))
}

func lexAndAssertEquality(input string, expected []par.TokenCode) AssertError {

	lexer := New(input)

	i := 0

	lexed := make([]par.Token, len(expected))

	for {
		tok := lexer.Lex()
		if tok.T == par.EOF {
			if i == len(expected) {

				return AssertError{NO_ERROR, ""}
			}
			return AssertError{LENGTH_INEQAULITY, prepareMessage(lexed, expected)}
		}

		if i >= len(expected) {
			return AssertError{LENGTH_INEQAULITY, prepareMessage(lexed, expected)}
		}

		lexed[i] = tok

		if expected[i] != tok.T {
			return AssertError{ELEMENT_INEQUALITY, prepareMessage(lexed, expected)}
		}

		i++
	}

}
func assert(input string, expected []par.TokenCode, expRes ASSERT_ERROR_CODE, t *testing.T) {
	result := lexAndAssertEquality(input, expected)

	if result.code != expRes {
		t.Fatal(result.msg)
	}
}

func assertNoError(input string, expected []par.TokenCode, t *testing.T) {
	assert(input, expected, NO_ERROR, t)
}

func TestSymbols(t *testing.T) {
	assertNoError(".", []par.TokenCode{par.DOT}, t)
	assertNoError("\"", []par.TokenCode{par.DOUBLEQUOTE}, t)
	assertNoError("'", []par.TokenCode{par.SINGLEQUOTE}, t)

	assertNoError("(", []par.TokenCode{par.LPAREN}, t)
	assertNoError(")", []par.TokenCode{par.RPAREN}, t)

	assertNoError("=", []par.TokenCode{par.EQ}, t)
	assertNoError("+", []par.TokenCode{par.PLUS}, t)
	assertNoError("-", []par.TokenCode{par.MINUS}, t)
	assertNoError("*", []par.TokenCode{par.MULTI}, t)

	assertNoError("\\", []par.TokenCode{par.BACKSLASH}, t)
	assertNoError("/", []par.TokenCode{par.SLASH}, t)
}

func TestIdent1(t *testing.T) {
	input := "Foo"

	expected := []par.TokenCode{
		par.IDENT,
	}

	assertNoError(input, expected, t)
}

func TestIdent2(t *testing.T) {
	input := "Foo2"

	expected := []par.TokenCode{
		par.IDENT,
	}

	assertNoError(input, expected, t)
}

func TestIdent3(t *testing.T) {
	input := "Foo+"

	expected := []par.TokenCode{
		par.IDENT,
		par.PLUS,
	}

	assertNoError(input, expected, t)
}

func TestSimpleBinaryExpression1(t *testing.T) {
	input := "2+2"

	expected := []par.TokenCode{
		par.NUMBER,
		par.PLUS,
		par.NUMBER,
	}

	assertNoError(input, expected, t)
}

func TestSimpleBinaryExpression2(t *testing.T) {
	input := "22+2"

	expected := []par.TokenCode{
		par.NUMBER,
		par.PLUS,
		par.NUMBER,
	}

	assertNoError(input, expected, t)
}

func TestSimpleBinaryExpression3(t *testing.T) {
	input := "22+111111"

	expected := []par.TokenCode{
		par.NUMBER,
		par.PLUS,
		par.NUMBER,
	}

	assertNoError(input, expected, t)
}

func TestSpacedTokens1(t *testing.T) {
	input := "2 2"

	expected := []par.TokenCode{
		par.NUMBER,
		par.NUMBER,
	}

	assertNoError(input, expected, t)
}

func TestSpacedTokens2(t *testing.T) {
	input := "Foo Baa"

	expected := []par.TokenCode{
		par.IDENT,
		par.IDENT,
	}

	assertNoError(input, expected, t)
}

func TestPlusSequence(t *testing.T) {
	input := "++"

	expected := []par.TokenCode{
		par.PLUS,
		par.PLUS,
	}

	assertNoError(input, expected, t)
}

func TestEmptyInput(t *testing.T) {
	input := ""

	expected := []par.TokenCode{
		par.NUMBER,
	}

	assert(input, expected, LENGTH_INEQAULITY, t)
}

func TestEmptyExpectation(t *testing.T) {
	input := "input"

	expected := []par.TokenCode{}

	assert(input, expected, LENGTH_INEQAULITY, t)
}
