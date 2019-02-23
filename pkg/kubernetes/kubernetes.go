package kubernetes

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"reflect"

	luar "github.com/geriBatai/gopher-luar"
	lua "github.com/yuin/gopher-lua"
)

var exports = map[string]lua.LGFunction{
	"Service":               bindResource(defaultService),
	"Secret":                bindResource(defaultSecret),
	"PersistentVolume":      bindResource(defaultPersistentVolume),
	"PersistentVolumeClaim": bindResource(defaultPersistentVolumeClaim),
	"Pod":                   bindResource(defaultPod),
	"ReplicationController": bindResource(defaultReplicationController),
	"ResourceQuota":         bindResource(defaultResourceQuota),
	"ServiceAccount":        bindResource(defaultServiceAccount),
	"Volume":                bindResource(defaultVolume),
	"DaemonSet":             bindResource(defaultDaemonSet),
	"Deployment":            bindResource(defaultDeployment),
	"ReplicaSet":            bindResource(defaultReplicaSet),
	"StatefulSet":           bindResource(defaultStatefulSet),
}

// Resource is an interface for generic Kubernetes
// object, similar to runtime.Object in Kubernetes code
type Resource interface {
	Clone() Resource
}

// defaultFunc return a default resource for any
// Kubernetes resource
type defaultFunc func(options map[string]interface{}) Resource

// Loader loads this module to lua
func Loader(l *lua.LState) int {
	mod := l.SetFuncs(l.NewTable(), exports)
	l.Push(mod)

	return 1
}

func copyResource(from, to Resource) Resource {
	buff := new(bytes.Buffer)
	enc := gob.NewEncoder(buff)
	dec := gob.NewDecoder(buff)
	enc.Encode(from)
	dec.Decode(to)
	return to
}

func bindResource(fn defaultFunc) lua.LGFunction {
	return func(L *lua.LState) int {
		var options = map[string]interface{}{}

		optionTb := L.OptTable(1, L.NewTable())
		optionTb.ForEach(func(key, value lua.LValue) {
			keystr, err := decodeLuaValue(key)
			if err != nil {
				fmt.Println(err.Error())
			}
			val, err := decodeLuaValue(value)
			options[keystr.(string)] = val
		})

		obj := fn(options)
		L.Push(luar.New(L, obj))
		return 1
	}
}

func decodeLuaValue(val lua.LValue) (interface{}, error) {
	switch v := val.(type) {
	case *lua.LNilType:
		return nil, nil
	case lua.LBool:
		return bool(v), nil
	case lua.LString:
		return string(v), nil
	case lua.LNumber:
		return float64(v), nil
	}
	return nil, fmt.Errorf("Cannot decode value %v of type %v", val, reflect.TypeOf(val))
}
