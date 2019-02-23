package kubernetes

import (
	v1 "k8s.io/api/core/v1"
)

// Volume is a wrapper around Kubernetes runtime.Object
// for lua
type Volume struct {
	*v1.Volume
}

// Clone returns a duplicate object. Used in lua as object::clone()
func (v *Volume) Clone() Resource {
	return copyResource(v, &Volume{})
}

func defaultVolume(options map[string]interface{}) Resource {
	return &Volume{
		Volume: &v1.Volume{},
	}
}
