package main

import (
	// "bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	// "reflect"
)

type BlockHeader struct {
	version     int32  // 4
	prev_block  []byte // 32
	merkle_root []byte // 32
	timestamp   uint32 // 4
	bits        uint32 // 4
	nonce       uint32 // 4
	// txn_count ?
	// txns ?
}

// type Transaction struct {
//   version uint32  // 4
//   inCount uint64  // variable
//   // inputs  //
// }

// but see also:
// https://stackoverflow.com/questions/28012952/golang-variable-length-array-in-struct-for-use-with-binary-read

// http://learnmeabitcoin.com/glossary/blkdat

// func readVarInt(f *os.File) []byte {
func readVarInt(f *os.File) uint64 {
	varintBuf := make([]byte, 1, 9)
	var rv uint64 = 0

	_, err2 := io.ReadFull(f, varintBuf[0:1])
	if err2 != nil {
		panic(err2)
	}

	iv := uint8(varintBuf[0])
	fmt.Printf("iv = %v\n", iv)
	// if uint8(value[0:1]) < 0xFD
	if iv < 0xFD {
		fmt.Println("Only one byte needed for this varint...")
		rv = uint64(iv)
		return rv
	}

	// another read is required...

	var readSize uint8 = 0
	if iv == 0xFD {
		// 0xfd followed by the length as uint16
		readSize = 2
	} else if iv == 0xFE {
		// 0xfe followed by the length as uint32
		readSize = 4
	} else if iv == 0xFF {
		// 0xff followed by the length as uint64
		readSize = 8
	}

	_, err2 = io.ReadFull(f, varintBuf[0:readSize])
	if err2 != nil {
		panic(err2)
	}

	return uint64(varintBuf[0])
}

func main() {
	// read.
	var buf [88]byte
	// var numTxBuf [9]byte

	f, err := os.Open("testdata/blk00028.dat")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err2 := io.ReadFull(f, buf[:])
	if err2 != nil {
		panic(err2)
	}

	var numTx = readVarInt(f)

	magic := binary.BigEndian.Uint32(buf[0:4])
	fmt.Printf("magic = %v\n", magic)

	size := binary.LittleEndian.Uint32(buf[4:8])
	fmt.Printf("size = %v\n", size)

	var blk_header BlockHeader
	// blk_header := buf[8:88]
	// fmt.Printf("blk_header = %v\n", blk_header)

	blk_header.version = int32(binary.LittleEndian.Uint32(buf[8:12]))
	blk_header.prev_block = buf[12:44]
	blk_header.merkle_root = buf[44:76]
	blk_header.timestamp = binary.LittleEndian.Uint32(buf[76:80])
	blk_header.bits = binary.LittleEndian.Uint32(buf[80:84])
	blk_header.nonce = binary.LittleEndian.Uint32(buf[84:88])

	fmt.Println("blk_header")
	fmt.Println("\t version:", blk_header.version)
	fmt.Println("\t prev_block:", blk_header.prev_block)
	fmt.Println("\t merkle_root:", blk_header.merkle_root)
	fmt.Println("\t timestamp:", blk_header.timestamp)
	fmt.Println("\t bits:", blk_header.bits)
	fmt.Println("\t nonce:", blk_header.nonce)

	fmt.Printf("\n")
	fmt.Printf("\t block has [%v] transaction(s):\n", numTx)

	// TODO: loop thru & read numTx transactions...

}

//  BlockHeader
//      version int32 // 4
//      prev_block []byte // 32
//      merkle_root []byte // 32
//      timestamp uint32 // 4
//      bits uint32 // 4
//      nonce uint32 // 4

// var p reflect.Value = reflect.ValueOf(&x)
// var v reflect.Value = p.Elem()
// v.SetFloat(7.1)
//
// fmt.Println("value: ", v)
// fmt.Println("kind is float64: ", v.Kind() == reflect.Float64)
