package snowflake

import (
	"errors"
	"net"
)

//DefaultGenerator default generator
var (
	DefaultGenerator *Node
)

// DefaultNode returns a new snowflake node by default nodeID
func DefaultNode() {
	nodeID, err := initNodeID()
	if err != nil {
		panic(err)
	}
	DefaultGenerator = NewNode(nodeID)
}

// Generate creates and returns a unique snowflake ID
func Generate() uint64 {
	return DefaultGenerator.Generate()
}

// GenerateID creates and returns a unique snowflake ID
func GenerateID() ID {
	return ID(DefaultGenerator.Generate())
}

// 节点进程编号(nodeID)最大限制是2^10，编号要满足(nodeID < 1024)。
// 1.针对IPV4:
// IP最大 255.255.255.255。而（255+255+255+255) < 1024。
// 因此采用IP段数值相加即可生成唯一的nodeID，不受IP位限制。
//
// 2.针对IPV6:
// IP最大 ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff
// 为了保证相加生成出的工程进程编号 < 1024,思路是将每个 Bit 位的后6位相加。
// IPV6 ：2^ 6 = 64。64 * 8 = 512 < 1024。
// 这样在一定程度上也可以满足nodeID不重复的问题。
// 使用这种 IP 生成工作进程编号的方法,必须保证IP段相加不能重复
// 例如254.255和255.254会重复
func initNodeID() (nodeID uint32, err error) {
	var (
		addrs   []net.Addr
		addr    net.Addr
		ipNet   *net.IPNet
		isIPNet bool
	)
	if addrs, err = net.InterfaceAddrs(); err != nil {
		return
	}
	for _, addr = range addrs {
		if ipNet, isIPNet = addr.(*net.IPNet); isIPNet && ipNet.IP.IsGlobalUnicast() {
			//IPV4
			if ip := ipNet.IP.To4(); ip != nil {
				for _, byteNum := range ip {
					nodeID += uint32(byteNum) & 0xFF
				}
				return
			}
			//IPV6
			if ip := ipNet.IP.To16(); ip != nil {
				for _, byteNum := range ip {
					nodeID += uint32(byteNum) & 0111111
				}
				return
			}
		}
	}
	err = errors.New("Bad LocalHost InetAddress, please check your network")
	return
}
