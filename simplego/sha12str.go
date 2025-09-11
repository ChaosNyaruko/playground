package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	fmt.Println(Sha1ToStr([]byte{0, 0, 0, 0, 82, 76, 102, 87, 72, 66, 15, 156, 104, 141, 139, 68, 3, 215, 227, 240}))
}

// Sha1ToStr sha1转string
func Sha1ToStr(s []byte) string {
	ins := make([]byte, 20)
	copy(ins, s)
	swapByte(ins)
	return hex.EncodeToString(ins)
}

func swapByte(s []byte) {
	for i := 0; i < 20; i += 4 {
		s[i], s[i+1], s[i+2], s[i+3] = s[i+3], s[i+2], s[i+1], s[i]
	}
}
