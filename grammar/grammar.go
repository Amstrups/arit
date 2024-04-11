package grammar

/*
-- (E)BNF syntax --
1: expr      ::= sum
2: expr      ::= product
3: expr      ::= value


4: sum       ::= sum addsub prodvalue
5: sum       ::= prodvalue addsub prodvalue
6: sum       ::= sum '-' parensum
7: sum       ::= prodvalue '-' parensum

8: addsub    ::= '+' | '-'
9: muldiv    ::= '*' | '/'

10: parensum  ::= '(' sum ')'

11: prodvalue ::= value
12: prodvalue ::= product

13: product   ::= prodterm muldiv value
14: product   ::= prodterm muldiv parensum
15: product   ::= prodterm '/' parenprod

16: prodterm  ::= prodvalue
17: prodterm  ::= parensum

18: parenprod ::= '(' product ')'

19: value     ::= number
*/

const (
	expr = iota
	sum
	product
	value
	addsub
	prodvalue
	parensum
	muldiv
	prodterm
	parenprod

	T_plus
	T_minus
	T_multi
	T_div
	T_lparen
	T_rparen
	T_number
)

type rule []int
type rules []rule

func possibilities(r int) rules {
	switch r {
	case expr:
		return rules{{sum}, {product}, {value}}
	case sum:
		return rules{{sum, addsub, prodvalue}, {prodvalue, addsub, prodvalue}}
	case addsub:
		return rules{{T_plus}, {T_minus}}
	case muldiv:
		return rules{{T_multi}, {T_div}}
	case parensum:
		return rules{{T_lparen, sum, T_rparen}}
	case prodvalue:
		return rules{{value}, {product}}
	case product:
		return rules{{prodterm, muldiv, value}, {prodterm, muldiv, parensum}, {prodterm, T_div, parenprod}}
	case prodterm:
		return rules{{prodvalue}, {parensum}}
	case parenprod:
		return rules{{T_lparen, product, T_rparen}}
	case value:
		return rules{{T_number}}
	default:
		return rules{}
	}
}
