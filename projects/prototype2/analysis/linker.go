package analysis

import (
	"prototype2/sexpr"
)

type node struct {
	astNode sexpr.Box
	attrs   sexpr.Box
}

func (n node) String() {

}
