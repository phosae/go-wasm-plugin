//go:build tinygo.wasm

// Code generated by protoc-gen-go-plugin. DO NOT EDIT.
// versions:
// 	protoc-gen-go-plugin v0.0.1
// 	protoc               v3.21.5
// source: test/import/proto/foo/foo.proto

package foo

import (
	context "context"
	wasm "github.com/knqyf263/go-plugin/wasm"
)

var foo Foo

func RegisterFoo(p Foo) {
	foo = p
}

//export foo_hello
func _foo_hello(ptr, size uint32) uint64 {
	b := wasm.PtrToByte(ptr, size)
	var req Request
	if err := req.UnmarshalVT(b); err != nil {
		return 0
	}
	response, err := foo.Hello(context.Background(), req)
	if err != nil {
		return 0
	}

	b, err = response.MarshalVT()
	if err != nil {
		return 0
	}
	ptr, size = wasm.ByteToPtr(b)
	return (uint64(ptr) << uint64(32)) | uint64(size)
}
