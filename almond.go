package almond

import (
	_ "embed"
	"fmt"

	"github.com/dop251/goja"
)

//go:embed almond.js
var almondJs string

var Program = goja.MustCompile("almond.js", almondJs, false)

type Module struct {
	*goja.Runtime
}

func Enable(vm *goja.Runtime) (*Module, error) {
	if _, err := vm.RunProgram(Program); err != nil {
		return nil, err
	}
	return New(vm), nil
}

func New(vm *goja.Runtime) *Module {
	return &Module{vm}
}

func (vm *Module) Define(m string, script string) error {
	if _, err := vm.RunString(fmt.Sprintf(`define.predef = "%s"`, m)); err != nil {
		return err
	}
	if _, err := vm.RunScript(m, script); err != nil {
		return err
	}
	if _, err := vm.RunString(`define.predef = null`); err != nil {
		return err
	}
	return nil
}

func (vm *Module) Require(m string) (goja.Value, error) {
	s := fmt.Sprintf(`requirejs("%s")`, m)
	return vm.RunString(s)
}
