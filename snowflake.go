// github.com/twitter/snowflake in golang

/*
+--------------------------------------------------------------------------+
| 1 Bit Unused | 41 Bit Timestamp |  10 Bit NodeID  |   12 Bit Sequence ID |
+--------------------------------------------------------------------------+

ID Format
By default, the ID format follows the original Twitter snowflake format.

The ID as a whole is a 63 bit integer stored in an int64
41 bits are used to store a timestamp with millisecond precision, using a custom epoch.
10 bits are used to store a node id - a range from 0 through 1023.
12 bits are used to store a sequence number - a range from 0 through 4095.

refer 从一次 Snowflake 异常说起
https://tech.meituan.com/2017/04/21/mt-leaf.html
https://segmentfault.com/a/1190000011282426
*/

package snowflake

import (
	"math/rand"
	"sync"
	"time"
)

var (
	//epoch 时间偏移量，从2019年6月16日零点开始
	epoch = time.Date(2019, 6, 16, 0, 0, 0, 0, time.Local)
)

const (
	//sequenceBits 自增量占用比特
	sequenceBits = 12
	//workeridBits 工作进程ID比特
	workeridBits = 5
	//datacenteridBits 数据中心ID比特
	datacenteridBits = 5
	//nodeidBits 节点ID比特
	nodeidBits = datacenteridBits + workeridBits
	//sequenceMask 自增量掩码（最大值）
	sequenceMask = -1 ^ (-1 << sequenceBits)
	//datacenteridLeftShiftBits 数据中心ID左移比特数（位数）
	datacenteridLeftShiftBits = workeridBits + sequenceBits
	//workeridLeftShiftBits 工作进程ID左移比特数（位数）
	workeridLeftShiftBits = sequenceBits
	//nodeidLeftShiftBits 节点ID左移比特数（位数）
	nodeidLeftShiftBits = datacenteridBits + workeridBits + sequenceBits
	//timestampLeftShiftBits 时间戳左移比特数（位数）
	timestampLeftShiftBits = nodeidLeftShiftBits
	//workeridMax 工作进程ID最大值
	workeridMax = -1 ^ (-1 << workeridBits)
	//datacenteridMax 数据中心ID最大值
	datacenteridMax = -1 ^ (-1 << datacenteridBits)
	//nodeidMax 节点ID最大值
	nodeidMax = -1 ^ (-1 << nodeidBits)
)

// Node a node struct for a snowflake generator
type Node struct {
	timestamp    int64
	datacenterID uint32
	workerID     uint32
	sequence     int
	lock         sync.Mutex
}

// NewNode returns a new snowflake node that can be used to generate snowflake
func NewNode(nodeID uint32) *Node {
	if nodeID >= nodeidMax {
		panic("Invalid nodeID")
	}
	datacenterID := nodeID >> datacenteridBits
	workerID := nodeID & (-1 ^ (-1 << workeridBits))
	return NewWorker(datacenterID, workerID)
}

// NewWorker returns a new snowflake node that can be used to generate snowflake
func NewWorker(datacenterID, workerID uint32) *Node {
	if datacenterID >= datacenteridMax {
		panic("Invalid datacenterID")
	}
	if workerID >= workeridMax {
		panic("Invalid workerID")
	}
	return &Node{datacenterID: datacenterID, workerID: workerID}
}

// Generate creates and returns a unique snowflake ID
// ID生成器,长度为64bit,从高位到低位依次为
// 1bit   符号位
// 41bits 时间偏移量从2019年6月16日零点到现在的毫秒数
// 10bits 节点工作进程ID
// 12bits 同一个毫秒内的自增量
func (n *Node) Generate() uint64 {
	n.lock.Lock()
	defer n.lock.Unlock()
	now := time.Since(epoch).Nanoseconds() / time.Millisecond.Nanoseconds()
	if now == n.timestamp {
		n.sequence = (n.sequence + 1) & sequenceMask
		if n.sequence == 0 {
			//保证当前时间大于最后时间。时间回退会导致产生重复id
			n.wait()
		}
	} else {
		n.initSequence()
	}
	//设置最后时间偏移量
	n.timestamp = now
	return n.id()
}

// GenerateID creates and returns a unique snowflake ID
func (n *Node) GenerateID() ID {
	return ID(n.Generate())
}

// 生成id
func (n *Node) id() uint64 {
	return uint64((n.timestamp << timestampLeftShiftBits)) |
		uint64((n.datacenterID << datacenteridLeftShiftBits)) |
		uint64((n.workerID << workeridLeftShiftBits)) |
		uint64(n.sequence)
}

// 避免在跨毫秒时序列号总是归0
func (n *Node) initSequence() {
	n.sequence = rand.New(rand.NewSource(time.Now().UnixNano())).Intn(10)
}

// 不停获得时间，直到大于最后时间
func (n *Node) wait() {
	for {
		now := time.Since(epoch).Nanoseconds() / time.Millisecond.Nanoseconds()
		if now > n.timestamp {
			break
		}
		if now == n.timestamp {
			time.Sleep(time.Millisecond)
			continue
		}
		time.Sleep(time.Duration((n.timestamp - now)) * time.Millisecond)
	}
}
