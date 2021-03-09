package types

import "syscall/js"

type VertexArray struct {
	js js.Value
}

func NewVertexArray(pointer js.Value) *VertexArray {
	return &VertexArray{
		js: pointer,
	}
}

func (vertexArray *VertexArray) GetJs() js.Value {
	return vertexArray.js
}
