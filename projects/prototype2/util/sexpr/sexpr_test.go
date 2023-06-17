package sexpr

import (
	"fmt"
	"strings"
	"testing"
)

func compare(list Box, expected string) bool {
	actual := list.String()
	return MinifySexpr(actual) == MinifySexpr(expected)
}

func TestConsCell(t *testing.T) {
	l := Cons(Box{1}, Box{2})
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
		(1 (2 (3 (4 (5 null)))))
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

	str := strings.Builder{}
	onEnter := func(node Box) {
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
	onExit := func(node Box) {
		switch node.Data.(type) {
		case string:
		case int:
		default:
			str.WriteByte(')')
		}
	}

	TraversePreorder(root, onEnter, onExit)
	println(PrettifySexpr(str.String()))
}
