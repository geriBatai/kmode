package kubernetes

import (
	luar "github.com/geriBatai/gopher-luar"
	lua "github.com/yuin/gopher-lua"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
)

var exports = map[string]lua.LGFunction{
	"Service":               newService,
	"Secret":                newSecret,
	"PersistentVolume":      newPersistentVolume,
	"PersistentVolumeClaim": newPersistentVolumeClaim,
	"Pod":                   newPod,
	"ReplicationController": newReplicationController,
	"ResourceQuota":         newResourceQuota,
	"ServiceAccount":        newServiceAccount,
	"Volume":                newVolume,
	"DaemonSet":             newDaemonSet,
	"Deployment":            newDeployment,
	"ReplicaSet":            newReplicaSet,
	"StatefulSet":           newStatefulSet,
}

// Loader loads this module to lua
func Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), exports)
	L.Push(mod)

	return 1
}

func newReplicationController(L *lua.LState) int {
	obj := &v1.ReplicationController{}
	L.Push(luar.New(L, obj))
	return 1
}

func newResourceQuota(L *lua.LState) int {
	obj := &v1.ResourceQuota{}
	L.Push(luar.New(L, obj))
	return 1
}

func newServiceAccount(L *lua.LState) int {
	obj := &v1.ServiceAccount{}
	L.Push(luar.New(L, obj))
	return 1
}

func newVolume(L *lua.LState) int {
	obj := &v1.Volume{}
	L.Push(luar.New(L, obj))
	return 1
}

func newDaemonSet(L *lua.LState) int {
	obj := &appsv1.DaemonSet{}
	L.Push(luar.New(L, obj))
	return 1
}

func newReplicaSet(L *lua.LState) int {
	obj := &appsv1.ReplicaSet{}
	L.Push(luar.New(L, obj))
	return 1
}

func newStatefulSet(L *lua.LState) int {
	obj := &appsv1.StatefulSet{}
	L.Push(luar.New(L, obj))
	return 1
}
