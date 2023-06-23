package sexpr

import (
	"fmt"
	"strings"
	"testing"
)

func compare(list Sexpr, expected string) bool {
	actual := list.String()
	return MinifySexpr(actual) == MinifySexpr(expected)
}

func TestConsCell(t *testing.T) {
	l := Cons(Sexpr{1}, Sexpr{2})
	expected := `
		(1 2)
	`
	if !compare(l, expected) {
		t.Fatalf("Sexpr\n%s\nIs malformed", PrettifySexpr(l.String()))
	}
}

func TestBasicSexpr(t *testing.T) {
	l := S(1, 2, 3, 4, 5)
	expected := `
		(1 (2 (3 (4 (5 nil)))))
	`
	if !compare(l, expected) {
		t.Fatalf("Sexpr\n%s\nIs malformed", PrettifySexpr(l.String()))
	}
}

func TestPreorder(t *testing.T) {
	root :=
		S("ProgramRoot",
			S("Function", "main",
				S("Body",
					S("+", 74, 99),
					S("+", 93, 88),
					S("+", 12, -70),
				)))
	expected := ` (ProgramRoot (Function main (Body (+ 74 99) (+ 93 88) (+ 12 -70)))) `

	str := strings.Builder{}
	onEnter := func(node Sexpr) {
		switch node.Data.(type) {
		case string:
			str.WriteString(node.Data.(string))
			str.WriteByte(' ')
		case int:
			str.WriteString(fmt.Sprint(node.Data.(int)))
			str.WriteByte(' ')
		default:
			str.WriteByte('(')
		}
	}
	onExit := func(node Sexpr) {
		switch node.Data.(type) {
		case string:
		case int:
		default:
			str.WriteByte(')')
		}
	}

	TraversePreorder(root, onEnter, onExit)
	if MinifySexpr(expected) != MinifySexpr(str.String()) {
		t.Fatalf("Not equal:\n%s\n%s", PrettifySexpr(str.String()), PrettifySexpr(expected))
	}
}

func TestEquality(t *testing.T) {
	cmp := func(lhs, rhs any) bool { return lhs == rhs }
	{
		lhs := Cons(Sexpr{"Cons"}, Sexpr{"Cell"})
		rhs := Cons(Sexpr{"Cons"}, Sexpr{"Cell"})
		if !Equals(lhs, rhs, cmp) {
			t.Fatalf("Equality test failed:\n%s\n%s", lhs.StringReadable(),
				rhs.StringReadable())
		}
	}

	{
		lhs := S(1, 2, 3)
		rhs := S(1, 2, 3)
		if !Equals(lhs, rhs, cmp) {
			t.Fatalf("Equality test failed:\n%s\n%s", lhs.StringReadable(),
				rhs.StringReadable())
		}
	}

	{
		lhs := S(1, S(2, "a"), 3)
		rhs := S(1, S(2, "a"), 3)
		if !Equals(lhs, rhs, cmp) {
			t.Fatalf("Equality test failed:\n%s\n%s", lhs.StringReadable(),
				rhs.StringReadable())
		}
	}

	{
		lhs := S(1, nil)
		rhs := S(nil, 1)
		if Equals(lhs, rhs, cmp) {
			t.Fatalf("Equality test failed:\n%s\n%s", lhs.StringReadable(),
				rhs.StringReadable())
		}
	}

}
