package main

import "fmt"

type MessageHeader struct {
	magic    uint32
	command  [12]byte
	length   uint32
	checksum uint32
	// payload []byte
}

type Block struct {
	version     int32
	prev_block  []byte
	merkle_root []byte
	timestamp   uint32
	bits        uint32
	nonce       uint32
	// txn_count ?
	// txns ?
}

// but see also:
// https://stackoverflow.com/questions/28012952/golang-variable-length-array-in-struct-for-use-with-binary-read

func main() {
	blk := Block{
		0x01,
		[]byte("0000000000000059424a547d9a1e9332ed7d4d68f7999b210d17677a3fab4dd4"),
		[]byte("4a81d33e5ea11a69b61e20004bae4933c00b846fbeadc888e55f0ff490bba888"),
		1511090782,
		0x196ff9a6,
		1647122739}
	fmt.Println(blk)
}

// fmt.Println(Vertex{1, 2})
