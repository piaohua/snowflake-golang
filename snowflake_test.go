package snowflake

import (
	"testing"
	"time"
)

func TestNewNode(t *testing.T) {
	node := NewNode(0)
	id := node.GenerateID()
	nodeID := id.Node()
	if nodeID != 0 {
		t.Fatal("invalid nodeID")
	}
	node = NewNode(nodeidMax)
	id = node.GenerateID()
	nodeID = id.Node()
	if nodeID != nodeidMax {
		t.Fatal("invalid nodeID")
	}
}

func TestNewWorker(t *testing.T) {
	node := NewWorker(1, 2)
	id := node.GenerateID()
	centerID := id.Center()
	if centerID != 1 {
		t.Fatal("Invalid datacenterID")
	}
	workerID := id.Worker()
	if workerID != 2 {
		t.Fatal("Invalid workderID")
	}
	nodeID := id.Node()
	// t.Log("nodeID = ", nodeID)
	if nodeID != ((1<<datacenteridBits)|2) {
		t.Fatal("Invalid nodeID")
	}

	node = &Node{datacenterID: datacenteridMax, workerID: workeridMax}
	id = node.GenerateID()
	centerID = id.Center()
	if centerID != datacenteridMax {
		t.Fatal("Invalid datacenterID")
	}
	workerID = id.Worker()
	if workerID != workeridMax {
		t.Fatal("Invalid workderID")
	}
	nodeID = id.Node()
	// t.Log("nodeID = ", nodeID)
	if nodeID != nodeidMax {
		t.Fatal("Invalid nodeID")
	}
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

func TestNodeMax(t *testing.T) {
	node := &Node{datacenterID: datacenteridMax, workerID: workeridMax}
	id := node.GenerateID()
	t.Log("id=", id)
	nodeId := id.Node()
	t.Log("nodeId=", nodeId)
	if nodeId != nodeidMax {
		t.Fatal("invalid nodeid")
	}
}
