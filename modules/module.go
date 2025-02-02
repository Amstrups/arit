package modules

import (
	"arit/modules/random"
	"fmt"
)

type Submodule interface {
	AddToModule(*Module)
	Parse([]string) (any, error)
	Name() string
	Description() string
}

type Module struct {
	submodules map[string]Submodule
}

func (m *Module) RegisterWithName(name string, sub Submodule) error {
	_, ok := m.submodules[name]
	if ok {
		return fmt.Errorf("module %s already exists", name)
	}

	m.submodules[name] = sub

	return nil
}

func (m *Module) Register(sub Submodule) error {
  return m.RegisterWithName(sub)
	_, ok := m.submodules[name]
	if ok {
		return fmt.Errorf("module %s already exists", name)
	}

	m.submodules[name] = sub

	return nil
}

func New() *Module {
	m := &Module{}
	m.
	return m

}
