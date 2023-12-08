package main

import "unsafe"

//go:wasm-module vm_hooks
func read_args(data uint32) {}

//go:wasm-module vm_hooks
func write_result(data uint32, len uint32) {}

//go:wasm-module vm_hooks
func memory_grow(data uint32) {}

func abort(message uint, filename uint, line uint32, column uint32) {
	return
}

//export mark_used
func mark_used() {
	memory_grow(0)
}

//export user_entrypoint
func user_entrypoint(len int32) int32 {
	input := make([]uint8, len)

	byteSize := uint32(unsafe.Sizeof(input[0])) * uint32(len)

	read_args(uint32(input[0]))
	write_result(uint32(input[0]), byteSize)
	return 0
}

func main() {}
