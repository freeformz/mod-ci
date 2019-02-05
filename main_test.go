package main

import (
	"testing"
	"time"
)

func TestThis(t *testing.T) {
	t.Log(t.Name())
}

func this() { time.Sleep(1 * time.Millisecond) }

func BenchmarkThis(b *testing.B) {
	for i := 0; i < b.N; i++ {
		this()
	}
}
