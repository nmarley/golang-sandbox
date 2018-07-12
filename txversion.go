package main

import "fmt"

func main() {
	// set old (signed 32-bit) version
	// var versionOld int32 = -987005808
	var versionOld int32 = 2

	// split into version and type fields (`type` is a reserved keyword in Go)
	var version = int16(versionOld & 0xffff)
	var ntype = int16((versionOld >> 16) & 0xffff)

	// print old & new
	fmt.Println("versionOld =", versionOld)
	fmt.Println("version =", version)
	fmt.Println("ntype =", ntype)
}

// this->nVersion = (int16_t) (n32bitVersion & 0xffff);
// this->nType = (int16_t) ((n32bitVersion >> 16) & 0xffff);
