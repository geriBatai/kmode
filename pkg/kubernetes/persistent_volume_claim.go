package kubernetes

import (
	"bytes"
	"encoding/gob"

	luar "github.com/geriBatai/gopher-luar"
	lua "github.com/yuin/gopher-lua"
	v1 "k8s.io/api/core/v1"
)

type PersistentVolumeClaim struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	*v1.PersistentVolumeClaim
}

func (s *PersistentVolumeClaim) Copy() *PersistentVolumeClaim {
	newobj := &PersistentVolumeClaim{}
	buff := new(bytes.Buffer)
	enc := gob.NewEncoder(buff)
	dec := gob.NewDecoder(buff)
	enc.Encode(s)
	dec.Decode(newobj)

	return newobj
}

func newPersistentVolumeClaim(L *lua.LState) int {
	obj := &PersistentVolumeClaim{
		Kind:                  "PersistentVolumeClaim",
		APIVersion:            "v1",
		PersistentVolumeClaim: &v1.PersistentVolumeClaim{},
	}
	L.Push(luar.New(L, obj))
	return 1
}
