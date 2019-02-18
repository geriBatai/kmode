package kubernetes

import (
	"bytes"
	"encoding/gob"

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

// KubernetesResource is an interface for generic Kubernetes
// object, similar to runtime.Object in Kubernetes code
type KubernetesResource interface {
	Copy() KubernetesResource
}

// defaultFunc return a default resource for any
// Kubernetes resource
type defaultFunc func() KubernetesResource

// Loader loads this module to lua
func Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), exports)
	L.Push(mod)

	return 1
}

func cloneResource(from, to KubernetesResource) KubernetesResource {
	buff := new(bytes.Buffer)
	enc := gob.NewEncoder(buff)
	dec := gob.NewDecoder(buff)
	enc.Encode(from)
	dec.Decode(to)
	return to
}

func bindResource(fn defaultFunc) lua.LGFunction {
	return func(L *lua.LState) int {
		// defaultValues := L.Get(-1)
		// obj := fn(defaultValues)
		obj := fn()
		L.Push(luar.New(L, obj))
		return 1
	}
}
