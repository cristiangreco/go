// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (openbsd && 386) || (openbsd && amd64) || (openbsd && arm64)
// +build openbsd,386 openbsd,amd64 openbsd,arm64

package runtime

import "unsafe"

// The X versions of syscall expect the libc call to return a 64-bit result.
// Otherwise (the non-X version) expects a 32-bit result.
// This distinction is required because an error is indicated by returning -1,
// and we need to know whether to check 32 or 64 bits of the result.
// (Some libc functions that return 32 bits put junk in the upper 32 bits of AX.)

//go:linkname syscall_syscall syscall.syscall
//go:nosplit
//go:cgo_unsafe_args
func syscall_syscall(fn, a1, a2, a3 uintptr) (r1, r2, err uintptr) {
	entersyscall()
	libcCall(unsafe.Pointer(funcPC(syscall)), unsafe.Pointer(&fn))
	exitsyscall()
	return
}
func syscall()

//go:linkname syscall_syscallX syscall.syscallX
//go:nosplit
//go:cgo_unsafe_args
func syscall_syscallX(fn, a1, a2, a3 uintptr) (r1, r2, err uintptr) {
	entersyscall()
	libcCall(unsafe.Pointer(funcPC(syscallX)), unsafe.Pointer(&fn))
	exitsyscall()
	return
}
func syscallX()

//go:linkname syscall_syscall6 syscall.syscall6
//go:nosplit
//go:cgo_unsafe_args
func syscall_syscall6(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2, err uintptr) {
	entersyscall()
	libcCall(unsafe.Pointer(funcPC(syscall6)), unsafe.Pointer(&fn))
	exitsyscall()
	return
}
func syscall6()

//go:linkname syscall_syscall6X syscall.syscall6X
//go:nosplit
//go:cgo_unsafe_args
func syscall_syscall6X(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2, err uintptr) {
	entersyscall()
	libcCall(unsafe.Pointer(funcPC(syscall6X)), unsafe.Pointer(&fn))
	exitsyscall()
	return
}
func syscall6X()

//go:linkname syscall_syscall10 syscall.syscall10
//go:nosplit
//go:cgo_unsafe_args
func syscall_syscall10(fn, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10 uintptr) (r1, r2, err uintptr) {
	entersyscall()
	libcCall(unsafe.Pointer(funcPC(syscall10)), unsafe.Pointer(&fn))
	exitsyscall()
	return
}
func syscall10()

//go:linkname syscall_syscall10X syscall.syscall10X
//go:nosplit
//go:cgo_unsafe_args
func syscall_syscall10X(fn, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10 uintptr) (r1, r2, err uintptr) {
	entersyscall()
	libcCall(unsafe.Pointer(funcPC(syscall10X)), unsafe.Pointer(&fn))
	exitsyscall()
	return
}
func syscall10X()

//go:linkname syscall_rawSyscall syscall.rawSyscall
//go:nosplit
//go:cgo_unsafe_args
func syscall_rawSyscall(fn, a1, a2, a3 uintptr) (r1, r2, err uintptr) {
	libcCall(unsafe.Pointer(funcPC(syscall)), unsafe.Pointer(&fn))
	return
}

//go:linkname syscall_rawSyscall6 syscall.rawSyscall6
//go:nosplit
//go:cgo_unsafe_args
func syscall_rawSyscall6(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2, err uintptr) {
	libcCall(unsafe.Pointer(funcPC(syscall6)), unsafe.Pointer(&fn))
	return
}

//go:linkname syscall_rawSyscall6X syscall.rawSyscall6X
//go:nosplit
//go:cgo_unsafe_args
func syscall_rawSyscall6X(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2, err uintptr) {
	libcCall(unsafe.Pointer(funcPC(syscall6X)), unsafe.Pointer(&fn))
	return
}

//go:linkname syscall_rawSyscall10X syscall.rawSyscall10X
//go:nosplit
//go:cgo_unsafe_args
func syscall_rawSyscall10X(fn, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10 uintptr) (r1, r2, err uintptr) {
	libcCall(unsafe.Pointer(funcPC(syscall10X)), unsafe.Pointer(&fn))
	return
}
