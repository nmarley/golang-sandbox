package main

import (
    // "bytes"
    "encoding/binary"
    "fmt"
    "os"
    "io"
)

type BlockHeader struct {
    version int32 // 4
    prev_block []byte // 32
    merkle_root []byte // 32
    timestamp uint32 // 4
    bits uint32 // 4
    nonce uint32 // 4
    // txn_count ?
    // txns ?
}

// but see also:
// https://stackoverflow.com/questions/28012952/golang-variable-length-array-in-struct-for-use-with-binary-read

// http://learnmeabitcoin.com/glossary/blkdat

func main() {
    // read.
    var buf [88]byte

    f, err := os.Open("blk00000.dat")
    if err != nil {
       panic(err)
    }
    defer f.Close()

    _, err2 := io.ReadFull(f, buf[:])
    if err2 != nil {
       panic(err2)
    }

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

}

//  BlockHeader
//      version int32 // 4
//      prev_block []byte // 32
//      merkle_root []byte // 32
//      timestamp uint32 // 4
//      bits uint32 // 4
//      nonce uint32 // 4
