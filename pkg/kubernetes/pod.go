package kubernetes

import (
	"bytes"
	"encoding/gob"

	luar "github.com/geriBatai/gopher-luar"
	lua "github.com/yuin/gopher-lua"
	v1 "k8s.io/api/core/v1"
)

type Pod struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	*v1.Pod
}

func (s *Pod) Copy() *Pod {
	newobj := &Pod{}
	buff := new(bytes.Buffer)
	enc := gob.NewEncoder(buff)
	dec := gob.NewDecoder(buff)
	enc.Encode(s)
	dec.Decode(newobj)

	return newobj
}

func newPod(L *lua.LState) int {
	obj := &Pod{
		Kind:       "Pod",
		APIVersion: "v1",
		Pod:        &v1.Pod{},
	}
	L.Push(luar.New(L, obj))
	return 1
}
