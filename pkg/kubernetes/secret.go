package kubernetes

import (
	"bytes"
	"encoding/gob"
	"fmt"

	luar "github.com/geriBatai/gopher-luar"
	lua "github.com/yuin/gopher-lua"
	v1 "k8s.io/api/core/v1"
	"k8s.io/kubernetes/pkg/kubectl/generate/versioned"
)

type Secret struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	*v1.Secret
}

func (s *Secret) Copy() *Secret {
	newobj := &Secret{}
	buff := new(bytes.Buffer)
	enc := gob.NewEncoder(buff)
	dec := gob.NewDecoder(buff)
	enc.Encode(s)
	dec.Decode(newobj)

	return newobj
}

func newSecret(L *lua.LState) int {
	generator := versioned.SecretGeneratorV1{}
	opts := map[string]interface{}{}
	opts["name"] = "secret"
	o, err := generator.Generate(opts)
	if err != nil {
		fmt.Printf("ERROR generating Secret resource: %v\n", err)
	}
	obj := &Secret{
		Kind:       "Secret",
		APIVersion: "v1",
		Secret:     o.(*v1.Secret),
	}
	L.Push(luar.New(L, obj))
	return 1
}
