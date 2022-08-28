//go:build tinygo.wasm

// Code generated by protoc-gen-go-plugin. DO NOT EDIT.
// versions:
// 	protoc-gen-go-plugin v0.1.0
// 	protoc               v3.21.5
// source: cat/cat.proto

package cat

import (
	context "context"
	wasm "github.com/knqyf263/go-plugin/wasm"
)

const FileCatPluginAPIVersion = 1

//export file_cat_api_version
func _file_cat_api_version() uint64 {
	return FileCatPluginAPIVersion
}

var fileCat FileCat

func RegisterFileCat(p FileCat) {
	fileCat = p
}

//export file_cat_cat
func _file_cat_cat(ptr, size uint32) uint64 {
	b := wasm.PtrToByte(ptr, size)
	var req FileCatRequest
	if err := req.UnmarshalVT(b); err != nil {
		return 0
	}
	response, err := fileCat.Cat(context.Background(), req)
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
