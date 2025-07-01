# 简介

将 [almond](https://github.com/requirejs/almond) 添加到 [goja](https://github.com/dop251/goja) 中, 以便使用 [AMD](https://github.com/amdjs/amdjs-api/wiki/AMD) 模块

# 为什么使用 AMD 模块

因为 AMD 模块简单, 可以控制插件的加载, 定制化更容易

# 如何使用

```go
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
}
```
