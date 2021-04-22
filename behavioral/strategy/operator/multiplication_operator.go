package operator

type Multiplication struct{}

func (Multiplication) Apply(lval, rval int) int {
	return lval * rval
}
