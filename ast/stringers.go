package ast

import (
	"fmt"
	"strconv"
)

func (cb *CodeBlock) String() string {
	return fmt.Sprintf("codeblock: %v", cb.Children); 
}

func (ls *LetStatement) String() string {
	return "let " + ls.Id + " = " + ls.Expression.String()
}

func (i *IntLiteral) String() string {
	return strconv.Itoa(i.Value, )
}