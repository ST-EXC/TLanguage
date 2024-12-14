package object

type Enviroment struct {
	store map[string]Object
	outer *Enviroment
}

func NewEnvironment() *Enviroment {
	s := make(map[string]Object)
	return &Enviroment{store: s, outer: nil}
}

func NewEnclosedEnvironment(outer *Enviroment) *Enviroment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func (e *Enviroment) Get(name string) (Object, bool, *Enviroment) {
	obj, ok := e.store[name]
	env := e
	if !ok && e.outer != nil {
		obj, ok, env = e.outer.Get(name)
	}
	return obj, ok, env
}

func (e *Enviroment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
