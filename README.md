# gothread

## Overview

Use consistent API to operate threads in Go language.

Supported OS:

Windows √

Linux √

MacOS √

## Usage

import this package:

```go
import "github.com/xuges/gothread"
```

Get current thread id:

```go
id := gothread.GetId()
fmt.Println("current thread id:", id)
```

Pin to current thread:

```go
gothread.Pin()
defer gothread.Unpin()
```

Run on thread stack:
```go
// run on current thread
gothread.Run(func() {
    println(gothread.GetId())
})

// run on new thread
go gothread.Run(func() {
    println(gothread.GetId())
})
```