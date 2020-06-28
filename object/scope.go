package object

func NewScope() Scope {
	scope := Scope{Variables: make(map[string]Object)}
	return scope
}

type Scope struct {
	Variables map[string]Object
}

func (s Scope) Get(name string) Object {
	return s.Variables[name]
}

func (s Scope) Set(name string, value Object) {
	s.Variables[name] = value
}

func (s Scope) IsDeclared(name string) bool {
	_, ok := s.Variables[name]
	return ok
}
