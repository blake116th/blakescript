package ast

func (bs *BlakeScript) String() string {
	return "Script Root"
}

func (ls *LetStatement) String() string {
	return "let " + ls.id
}
