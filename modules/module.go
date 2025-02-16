package modules

import (
	subs "arit/modules/submodules"
	"fmt"
)

type Submodule interface {
	Parse([]string) (any, error)
	Name() string
	Keys() []string
	Description() string
}

type Module struct {
	Submodules map[string]Submodule
}

func (m *Module) Register(sub Submodule) error {
	name := sub.Name()
	_, ok := m.Submodules[name]
	if ok {
		return fmt.Errorf("module %s already exists", name)
	}

	m.Submodules[name] = sub

	for _, k := range sub.Keys() {
		if k == name {
			continue
		}
		m.Submodules[k] = sub
	}

	return nil
}

func Full() Module {
	m := Module{
		Submodules: map[string]Submodule{},
	}

	m.Register(&subs.Random{})
	m.Register(&subs.Prime{})

	return m

}
