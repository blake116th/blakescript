package ast

import (
	"fmt"
	"github.com/Appleby43/blakescript/vm"
)

type Statement interface {
	fmt.Stringer
	Execute(env vm.Env, scope *CodeBlock)
}

//A Codeblock is the basic unit of scope in blakescript 
type CodeBlock struct {
	parent *CodeBlock // todo use to ascend scope
	table vm.ScopeTable
	Children []Statement
}

func (cb *CodeBlock) Execute(env vm.Env, scope *CodeBlock) {
	for _, child := range cb.Children {
		child.Execute(env, cb)
	}
}

type Expression interface {
	fmt.Stringer
	Evaluate(env vm.Env, scope *CodeBlock) int
}

type IntLiteral struct {
	Expression
	Value int
}

func (i *IntLiteral) Evaluate(env vm.Env, scope *CodeBlock) int {
	return i.Value
}

type IdentifierExpression struct {
	Id string
}

func (i *IdentifierExpression) Evaluate(env vm.Env, scope *CodeBlock) int {
	heapId := scope.table[i.Id]
	return env.Heap[heapId]
}

type LetStatement struct {
	Statement
	Id string
	Expression Expression
}

func (l *LetStatement) Execute(env vm.Env, scope *CodeBlock) {
	heapId := env.MakeHeapId()

	env.Heap[heapId] = l.Expression.Evaluate(env, scope);
	scope.table[l.Id] = heapId;
}