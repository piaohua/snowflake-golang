# snowflake

    This sample is a very simple Twitter snowflake generator in [Go](https://golang.org/)

## ID Format
By default, the ID format follows the original Twitter snowflake format.
* The ID as a whole is a 63 bit integer stored in an int64
* 41 bits are used to store a timestamp with millisecond precision, using a custom epoch.
* 10 bits are used to store a node id - a range from 0 through 1023.
* 12 bits are used to store a sequence number - a range from 0 through 4095.

### How it Works.
Each time you generate an ID, it works, like this.
* A timestamp with millisecond precision is stored using 41 bits of the ID.
* Then the NodeID is added in subsequent bits.
* Then the Sequence Number is added, starting at 0 and incrementing for each ID generated in the same millisecond. If you generate enough IDs in the same millisecond that the sequence would roll over or overfill then the generate function will pause until the next millisecond.

The default Twitter format shown below.
```
+--------------------------------------------------------------------------+
| 1 Bit Unused | 41 Bit Timestamp |  10 Bit NodeID  |   12 Bit Sequence ID |
+--------------------------------------------------------------------------+
```

Using the default settings, this allows for 4096 unique IDs to be generated every millisecond, per Node ID.

## Installation

```sh
go get github.com/piaohua/snowflake-golang
```

## Usage:

```go
package main

import (
	"fmt"

	snowflake "github.com/piaohua/snowflake-golang"
)

func main() {

	// Create a new Node with a Node number of 1
	node := snowflake.NewNode(1)

	// Generate a snowflake ID.
	id := node.GenerateID()

	// Create a new Node with a datacenter and worker number of 1
	//node := snowflake.NewWorker(1, 1)
	//id := node.GenerateID()

	// Create a new Node with localhost
	//snowflake.DefaultNode()
    //id := snowflake.GenerateID()

	// Print out the ID in a few different ways.
	fmt.Printf("Uint64 ID: %d\n", id.Uint64())
	fmt.Printf("String ID: %s\n", id.String())
	fmt.Printf("Base2  ID: %s\n", id.Base2())
	fmt.Printf("Base64 ID: %s\n", id.Base64())

	// Print out the ID's timestamp
	fmt.Printf("ID Time  : %d\n", id.Time())

	// Print out the ID's node number
	fmt.Printf("ID Node  : %d\n", id.Node())

	// Print out the ID's sequence number
	fmt.Printf("ID Sequence  : %d\n", id.Sequence())
}
```

## References:
*  https://github.com/bwmarrin/snowflake
*  https://tech.meituan.com/2017/04/21/mt-leaf.html
*  https://segmentfault.com/a/1190000011282426
