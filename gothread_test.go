package gothread

import (
	"testing"
	"time"
)

func Test_GetId(t *testing.T) {
	Run(func() {
		println("thread id:", GetId())
	})
}

func Test_MultiThread(t *testing.T) {
	go Run(func() {
		println("1 thread id:", GetId())
	})

	go Run(func() {
		println("2 thread id:", GetId())
	})

	Run(func() {
		println("3 thread id:", GetId())
	})
	time.Sleep(1 * time.Second)
}

func Test_CloseThread(t *testing.T) {
	go println("1 thread id:", GetId())
	println("2 thread id:", GetId())
	go func() {
		println("3 thread id:", GetId())
		Close()
	}()
	go func() {
		time.Sleep(1 * time.Second)
		println("4 thread id:", GetId())
	}()
	time.Sleep(2 * time.Second)
}
