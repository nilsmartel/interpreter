package execution

import "interpreter/value"

type Callable interface {
	call(env *Env, args []value.Object) (value.Object, error)
}
