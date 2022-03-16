package ast

import (
	"fmt"
	"github.com/Appleby43/blakescript/vm"
)

type Node interface {
	fmt.Stringer
	Execute(env vm.Env)
}

type BlakeScript struct {
	children []Node
}

func (bs *BlakeScript) Execute(env vm.Env) {
	for _, child := range bs.children {
		child.Execute(env)
	}
}

type Expression interface {
	Node
	Evaluate() int
}

type LetStatement struct {
	Node
	id string
	Expression Expression
}

func (let *LetStatement) Execute(env vm.Env) {
	id := env.MakeHeapId()
	//todo place heap id in scoped table
	env.Heap[id] = let.Expression.Evaluate();
}