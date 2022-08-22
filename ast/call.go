package ast

// (do-stuff x y)
type Call struct {
	Function  Expression
	Arguments Expression
}
