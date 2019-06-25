package snowflake

import (
	"testing"
	"time"
)

func TestNewNode(t *testing.T) {
	NewNode(0)
	NewNode(1024)
}

func TestNewWorker(t *testing.T) {
	NewWorker(0, 0)
	NewWorker(32, 32)
}

func TestGenerate(t *testing.T) {
	start := time.Now()
	defer func() {
		t.Logf("elapsed time: %s\n", time.Since(start))
	}()
	node := NewNode(0)
	var x, y uint64
	for i := 0; i < 4096; i++ {
		y = node.Generate()
		if x == y {
			t.Errorf("x(%d) & y(%d) are the same", x, y)
		}
		x = y
	}
}

func TestDefaultNode(t *testing.T) {
	start := time.Now()
	defer func() {
		t.Logf("elapsed time: %s\n", time.Since(start))
	}()
	DefaultNode()
	var x, y uint64
	for i := 0; i < 1000000; i++ {
		y = Generate()
		if x == y {
			t.Errorf("x(%d) & y(%d) are the same", x, y)
		}
		x = y
	}
}
