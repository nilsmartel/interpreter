package ast

// ((if x y do-stuff) x y)
type Call struct {
	Function  Expression
	Arguments []Expression
}

// just like call, but it's a little more direct
// (cosine x y)
type NamedCall struct {
	Function  string
	Arguments []Expression
}
