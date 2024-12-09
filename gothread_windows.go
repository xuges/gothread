package gothread

import "golang.org/x/sys/windows"

var getThreadId = windows.GetCurrentThreadId
