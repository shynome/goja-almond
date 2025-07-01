package almond_test

import (
	_ "embed"
	"testing"

	"github.com/dop251/goja"
	"github.com/shynome/err0/try"
	almond "github.com/shynome/goja-almond"
)

//go:embed test.js
var tJs string

func TestAlmond(t *testing.T) {
	vm := goja.New()
	mod := try.To1(almond.Enable(vm))
	try.To(mod.Define("test.js", tJs))
	v := try.To1(mod.Require("test.js"))
	if v := v.String(); v != "test" {
		t.Errorf("want test, but got %s", v)
		return
	}
	t.Log(v)
	//
	_, err := vm.RunScript("test2.js", tJs)
	if err == nil {
		t.Error("should be wrong, but err is nil")
	}
	errStr := err.Error()
	t.Log(errStr)
	//
	try.To(mod.Define("test3.js", tJs))
	v3 := try.To1(mod.Require("test3.js"))
	if v := v3.String(); v != "test" {
		t.Errorf("want test, but got %s", v)
		return
	}
}
