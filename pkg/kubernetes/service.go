package kubernetes

import (
	"bytes"
	"encoding/gob"
	"fmt"

	luar "github.com/geriBatai/gopher-luar"
	lua "github.com/yuin/gopher-lua"
	v1 "k8s.io/api/core/v1"

	//  "github.com/kubernetes/kubernetes/pkg/kubectl/generate/versioned"
	"k8s.io/kubernetes/pkg/kubectl/generate/versioned"
)

type Service struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	*v1.Service
}

func (s *Service) Copy() *Service {
	newobj := &Service{}
	buff := new(bytes.Buffer)
	enc := gob.NewEncoder(buff)
	dec := gob.NewDecoder(buff)
	enc.Encode(s)
	dec.Decode(newobj)

	return newobj
}

func newService(L *lua.LState) int {
	// "github.com/kubernetes/kubernetes/pkg/kubectl/generate/versioned"
	generator := versioned.ServiceGeneratorV1{}
	opts := map[string]interface{}{}
	opts["default-name"] = "svc"
	opts["selector"] = "name=svc"
	opts["port"] = "80"

	s, err := generator.Generate(opts)
	if err != nil {
		fmt.Printf("ERROR generating Service resource: %v\n", err)
	}

	obj := &Service{
		Kind:       "Service",
		APIVersion: "v1",
		Service:    s.(*v1.Service),
	}
	L.Push(luar.New(L, obj))
	return 1
}
