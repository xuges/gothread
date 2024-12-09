package gothread

import "golang.org/x/sys/unix"

var getThreadId = unix.Gettid
