package gothread

import (
	"runtime"
)

// Pin locks the current goroutine to OS thread
func Pin() {
	runtime.LockOSThread()
}

// Unpin unlocks the pined OS thread and current goroutine
func Unpin() {
	runtime.UnlockOSThread()
}

// Run runs the task on OS thread
func Run(task func()) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	task()
}

// Close the OS thread
//
// NOTE: must call on non-pined goroutine
func Close() {
	runtime.LockOSThread()
}

// GetId returns the current OS thread id
func GetId() int {
	return int(getThreadId())
}
