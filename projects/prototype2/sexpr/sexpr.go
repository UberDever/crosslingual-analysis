package sexpr

import (
	"fmt"
	"strings"
	"unicode"
)

type Sexpr struct {
	Data any
}

func (v Sexpr) IsAtom() bool {
	switch v.Data.(type) {
	case Sexpr:
		return false
	case cell:
		return false
	default:
		return true
	}
}

func (v Sexpr) IsNil() bool {
	return v.Data == nil
}

func (v Sexpr) StringReadable() string {
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

	TraversePreorder(v, onEnter, onExit)
	return PrettifySexpr(str.String())
}

func (v Sexpr) String() string {
	toString := func(v Sexpr) string {
		switch v.Data.(type) {
		case int:
			return fmt.Sprint(v.Data)
		case string:
			return fmt.Sprintf("'%s'", v.Data.(string))
		case nil:
			return "nil"
		}
		return ""
	}
	s := strings.Builder{}
	vStack := make([]Sexpr, 0, 64)
	vStack = append(vStack, v)
	branchesWalked := make([]int, 0, 32)
	branchesWalked = append(branchesWalked, 0)

	for {
		if len(vStack) == 0 {
			break
		}

		top := vStack[len(vStack)-1]
		vStack = vStack[:len(vStack)-1]
		branchesWalked[len(branchesWalked)-1]++

		switch top.Data.(type) {
		case cell:
			c := top.Data.(cell)
			s.WriteByte('(')
			vStack = append(vStack, c.rhs)
			vStack = append(vStack, c.lhs)
			branchesWalked = append(branchesWalked, 0)
		default:
			s.WriteString(toString(top))
			s.WriteByte(' ')
		}

		for branchesWalked[len(branchesWalked)-1] == 2 {
			s.WriteByte(')')
			branchesWalked = branchesWalked[:len(branchesWalked)-1]
			if len(branchesWalked) == 0 {
				break
			}
		}

	}
	return s.String()
}

type cell struct {
	lhs Sexpr
	rhs Sexpr
}

func S(list ...any) Sexpr {
	if len(list) == 0 {
		return Sexpr{nil}
	}
	box := Sexpr{}
	switch list[0].(type) {
	case Sexpr:
		box = list[0].(Sexpr)
	default:
		box.Data = list[0]
	}
	return Cons(box, S(list[1:]...))
}

func Cons(lhs any, rhs any) Sexpr {
	var l, r Sexpr
	switch lhs.(type) {
	case Sexpr:
		l = lhs.(Sexpr)
	default:
		l = Sexpr{lhs}
	}
	switch rhs.(type) {
	case Sexpr:
		r = rhs.(Sexpr)
	default:
		r = Sexpr{rhs}
	}
	return Sexpr{cell{lhs: l, rhs: r}}
}

func Car(v Sexpr) Sexpr {
	switch v.Data.(type) {
	case cell:
		unboxed := v.Data.(cell)
		return unboxed.lhs
	default:
		return Sexpr{nil}
	}
}

func Cdr(v Sexpr) Sexpr {
	switch v.Data.(type) {
	case cell:
		unboxed := v.Data.(cell)
		return unboxed.rhs
	default:
		return Sexpr{nil}
	}
}

func Equals(lhs Sexpr, rhs Sexpr, cmp func(any, any) bool) bool {
	if lhs.IsAtom() || rhs.IsAtom() {
		return cmp(lhs.Data, rhs.Data)
	}

	return Equals(Car(lhs), Car(rhs), cmp) &&
		Equals(Cdr(lhs), Cdr(rhs), cmp)
}

func PrettifySexpr(sexpr string) string {
	formatted := strings.Builder{}
	depth := -1
	for i := range sexpr {
		if sexpr[i] == '(' {
			depth++
			formatted.WriteByte('\n')
			for j := 0; j < depth; j++ {
				formatted.WriteString("    ")
			}
			formatted.WriteByte('(')
		} else if sexpr[i] == ')' {
			depth--
			formatted.WriteByte(')')
		} else if unicode.IsSpace(rune(sexpr[i])) {
			if i < len(sexpr)-1 && sexpr[i+1] != ')' {
				formatted.WriteByte(sexpr[i])
			}
		} else {
			formatted.WriteByte(sexpr[i])
		}
	}
	return formatted.String()
}

func MinifySexpr(s string) string {
	formatted := strings.Builder{}
	skipWS := func(i int) (int, bool) {
		wasSpace := false
		for s[i] == ' ' || s[i] == '\n' || s[i] == '\t' {
			wasSpace = true
			i++
			if i >= len(s) {
				break
			}
		}
		return i, wasSpace
	}

	for i := 0; i < len(s); i++ {
		j, wasSpace := skipWS(i)
		if j >= len(s) {
			break
		}
		i = j
		if wasSpace {
			if s[i] != '(' && s[i] != ')' {
				formatted.WriteByte(' ')
			}
		}
		formatted.WriteByte(s[i])
	}
	return formatted.String()
}

type Action func(Sexpr)

func TraversePreorder(root Sexpr, onEnter Action, onExit Action) {
	traversePreorderRec(onEnter, onExit, root)
}

func traversePreorderRec(onEnter Action, onExit Action, cur Sexpr) {
	if cur.Data == nil {
		return
	}

	onEnter(cur)
	defer onExit(cur)
	children := cur
	for c := Car(children); c.Data != nil; c = Car(children) {
		children = Cdr(children)
		traversePreorderRec(onEnter, onExit, c)
	}
}

func TraversePostorder(root Sexpr, onEnter Action) {
	traversePostorderRec(onEnter, root)
}

func traversePostorderRec(onEnter Action, cur Sexpr) {
	c := Car(cur)
	if c.Data == nil {
		return
	}

	traversePostorderRec(onEnter, Cdr(cur))
	traversePostorderRec(onEnter, c)
	onEnter(c)
}
