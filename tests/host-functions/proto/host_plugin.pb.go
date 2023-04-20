//go:build tinygo.wasm

// Code generated by protoc-gen-go-plugin. DO NOT EDIT.
// versions:
// 	protoc-gen-go-plugin v0.1.0
// 	protoc               v3.21.12
// source: tests/host-functions/proto/host.proto

package proto

import (
	context "context"
	wasm "github.com/knqyf263/go-plugin/wasm"
	_ "unsafe"
)

const GreeterPluginAPIVersion = 1

//export greeter_api_version
func _greeter_api_version() uint64 {
	return GreeterPluginAPIVersion
}

var greeter Greeter

func RegisterGreeter(p Greeter) {
	greeter = p
}

//export greeter_greet
func _greeter_greet(ptr, size uint32) uint64 {
	b := wasm.PtrToByte(ptr, size)
	req := new(GreetRequest)
	if err := req.UnmarshalVT(b); err != nil {
		return 0
	}
	response, err := greeter.Greet(context.Background(), req)
	if err != nil {
		ptr, size = wasm.ByteToPtr([]byte(err.Error()))
		return (uint64(ptr) << uint64(32)) | uint64(size) |
			// Indicate that this is the error string by setting the 32-th bit, assuming that
			// no data exceeds 31-bit size (2 GiB).
			(1 << 31)
	}

	b, err = response.MarshalVT()
	if err != nil {
		return 0
	}
	ptr, size = wasm.ByteToPtr(b)
	return (uint64(ptr) << uint64(32)) | uint64(size)
}

type hostFunctions struct{}

func NewHostFunctions() HostFunctions {
	return hostFunctions{}
}

//go:wasm-module env
//export parse_json
//go:linkname _parse_json
func _parse_json(ptr uint32, size uint32) uint64

func (h hostFunctions) ParseJson(ctx context.Context, request *ParseJsonRequest) (*ParseJsonResponse, error) {
	buf, err := request.MarshalVT()
	if err != nil {
		return nil, err
	}
	ptr, size := wasm.ByteToPtr(buf)
	ptrSize := _parse_json(ptr, size)
	wasm.FreePtr(ptr)

	ptr = uint32(ptrSize >> 32)
	size = uint32(ptrSize)
	buf = wasm.PtrToByte(ptr, size)

	response := new(ParseJsonResponse)
	if err = response.UnmarshalVT(buf); err != nil {
		return nil, err
	}
	return response, nil
}
