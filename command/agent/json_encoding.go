package agent

import (
	"reflect"

	"github.com/hashicorp/go-msgpack/codec"

	"github.com/hashicorp/nomad/nomad/structs"
)

// Special encoding for structs.Node, to perform the following:
// 1. ensure that Node.SecretID is zeroed out
// 2. provide backwards compatibility for the following fields:
//   * Node.Drain
type nodesExt struct{}

// ConvertExt converts a structs.Node to an anonymous struct with the extra field, Drain
func (n nodesExt) ConvertExt(v interface{}) interface{} {
	node := v.(*structs.Node)
	copy := node.Copy()
	copy.SecretID = ""
	return struct {
		*structs.Node
		Drain bool
	}{
		Node:  copy,
		Drain: node.DrainStrategy != nil,
	}
}

// UpdateExt is not used
func (n nodesExt) UpdateExt(_ interface{}, _ interface{}) {}

func registerExtensions(h *codec.JsonHandle) *codec.JsonHandle {
	h.SetInterfaceExt(reflect.TypeOf(structs.Node{}), 1, nodesExt{})
	return h
}
