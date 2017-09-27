package main

import (
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
	"crypto/rand"
)

var escaper = strings.NewReplacer("9", "99", "-", "90", "_", "91")
var unescaper = strings.NewReplacer("99", "9", "90", "-", "91", "_")

func main() {
	uuid := []byte{0x12, 0x3e, 0x45, 0x67, 0xe8, 0x9b, 0x12, 0xd3, 0xa4, 0x56, 0x42, 0x66, 0x55, 0x44, 0x00, 0x00}
	fmt.Printf("%x\n", uuid)
	fmt.Println(hex.EncodeToString(uuid))

	fmt.Println(base32.StdEncoding.EncodeToString(uuid))

	fmt.Println(base64.RawURLEncoding.EncodeToString(uuid))

	fmt.Println(escaper.Replace(base64.RawURLEncoding.EncodeToString(uuid)))

	fmt.Println("Verify decoding:")
	s := escaper.Replace(base64.RawURLEncoding.EncodeToString(uuid))
	dec, err := base64.RawURLEncoding.DecodeString(unescaper.Replace(s))
	fmt.Printf("%x, %v\n", dec, err)

	key := uuidFunc()
	fmt.Println(key)
	key = uuidFunc()
	fmt.Println(key)
}

func uuidFunc() string {
    b := make([]byte, 16)
	rand.Read(b)
    b[6] = (b[6] & 0x0f) | 0x40
    b[8] = (b[8] & 0x3f) | 0x80
    return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
