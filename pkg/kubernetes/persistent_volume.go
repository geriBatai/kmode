package kubernetes

import (
	"bytes"
	"encoding/gob"

	luar "github.com/geriBatai/gopher-luar"
	lua "github.com/yuin/gopher-lua"
	v1 "k8s.io/api/core/v1"
)

type PersistentVolume struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	*v1.PersistentVolume
}

func (s *PersistentVolume) Copy() *PersistentVolume {
	newobj := &PersistentVolume{}
	buff := new(bytes.Buffer)
	enc := gob.NewEncoder(buff)
	dec := gob.NewDecoder(buff)
	enc.Encode(s)
	dec.Decode(newobj)

	return newobj
}

func newPersistentVolume(L *lua.LState) int {
	obj := &PersistentVolume{
		Kind:             "PersistentVolume",
		APIVersion:       "v1",
		PersistentVolume: &v1.PersistentVolume{},
	}
	L.Push(luar.New(L, obj))
	return 1
}
