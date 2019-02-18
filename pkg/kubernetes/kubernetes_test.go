package kubernetes_test

import (
	"github.com/geriBatai/kmode/pkg/kubernetes"
	lua "github.com/yuin/gopher-lua"
)

func buildLuaObject(o string) (*lua.LUserData, error) {
	L := lua.NewState()
	defer L.Close()
	L.PreloadModule("kubernetes", kubernetes.Loader)
	code := `local kubernetes = require("kubernetes"); data = ` + o
	err := L.DoString(code)
	if err != nil {
		return nil, err
	}
	return L.GetGlobal("data").(*lua.LUserData), nil
}
