package main

import "fmt"

type Just[T any] struct {
	a T
}

func (Just[T]) MaybeVariant() {}

type Nothing struct{}

func (Nothing) MaybeVariant() {}

type Maybe interface{ MaybeVariant() }

func foo[T any](v Maybe) {
	switch val := v.(type) {
	case Just[T]:
		fmt.Println(val)
	case Nothing:
		fmt.Println(val)
	default:
		panic("Unreachable")
	}
}

type Branch[T any] struct {
	lhs, rhs Tree
}

func (Branch[T]) TreeVariant() {}

type Leaf[T any] struct {
	a T
}

func (Leaf[T]) TreeVariant() {}

type Tree interface{ TreeVariant() }

func printTree[T any](t Tree) {
	switch v := t.(type) {
	case Branch[T]:
		printTree[T](v.lhs)
		printTree[T](v.rhs)
	case Leaf[T]:
		fmt.Println(v.a)
	default:
		panic("Unreachable")
	}
}

func main() {
	a := Just[int]{5}
	b := Nothing{}
	foo[int](a)
	foo[struct{}](b)

	tree := Branch[int]{
		Leaf[int]{1},
		Branch[int]{
			Branch[int]{
				Leaf[int]{16},
				Leaf[int]{28},
			},
			Leaf[int]{10},
		},
	}
	printTree[int](tree)
}
