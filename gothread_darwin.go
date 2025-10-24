package gothread

import (
	_ "syscall"
	_ "unsafe"
)

//go:linkname runtime_rawSyscall syscall.rawSyscall
func runtime_rawSyscall(fn, a1, a2, a3 uintptr) (r1, r2, err uintptr)

//go:linkname runtime_syscall syscall.rawSyscall
func runtime_syscall(fn, a1, a2, a3 uintptr) (r1, r2, err uintptr)

//go:cgo_import_dynamic libpthread_pthread_mach_thread_np pthread_mach_thread_np "/usr/lib/libpthread.dylib"
var libpthread_pthread_mach_thread_np_trampoline_addr uintptr

func pthread_mach_thread_np(t uintptr) uintptr {
	ret, _, _ := runtime_syscall(libpthread_pthread_mach_thread_np_trampoline_addr, t, 0, 0)
	return ret
}

//go:cgo_import_dynamic libpthread_pthread_self pthread_self "/usr/lib/libpthread.dylib"
var libpthread_pthread_self_trampoline_addr uintptr

func pthread_self() uintptr {
	ret, _, _ := runtime_syscall(libpthread_pthread_self_trampoline_addr, 0, 0, 0)
	return ret
}

func getThreadId() int {
	return int(pthread_mach_thread_np(pthread_self()))
}
