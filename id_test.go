package snowflake

import (
	"testing"
	"time"
)

func TestID(t *testing.T) {
	start := time.Now()
	defer func() {
		t.Logf("elapsed time: %s\n", time.Since(start))
	}()
	DefaultNode()
	id := GenerateID()
	t.Logf("id %d", id.Uint64())
	t.Log(id.Base2())
	t.Log(id.Time(), id.Node(), id.Sequence())
}

func TestParse(t *testing.T) {
	//time.Now().UnixNano() = 1560676238541353000
	//ts = 61838541
	//1560676238541000000
	id := ID(259369641758724)
	t.Log(id.Time(), id.Node(), id.Sequence())
	if id.Base2() != "111010111110010100110011010111001101000000000100" {
		t.Fatal("failed")
	}
	if id.Node() != 461 {
		t.Fatal("failed")
	}
	if id.Sequence() != 4 {
		t.Fatal("failed")
	}
	//2019-06-16 17:10:38.541 +0000 UTC
	nano := id.Time() / 1000000
	t.Log(time.Unix((nano / 1000), ((nano % 1000) * 1000000)))
}

func TestMaxID(t *testing.T) {
	id := ID((1 << 63) - 1)
	t.Log(id, id.Time(), id.Node(), id.Sequence())
	if id.Node() != 1023 {
		t.Fatal("failed")
	}
	if id.Sequence() != 4095 {
		t.Fatal("failed")
	}
	//2089-02-19 15:47:35.551 +0000 UTC
	nano := id.Time() / 1000000
	t.Log(time.Unix((nano / 1000), ((nano % 1000) * 1000000)))
}

func TestDefaultNodeID(t *testing.T) {
	start := time.Now()
	defer func() {
		t.Logf("elapsed time: %s\n", time.Since(start))
	}()
	nodeID, err := initNodeID()
	if err != nil {
		t.Fatal(err)
	}
	//192.168.0.101 == 461
	if nodeID != 461 {
		t.Fatal("failed")
	}
}
