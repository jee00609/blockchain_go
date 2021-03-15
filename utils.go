package main

import (
	"bytes"
	"encoding/binary"
	"log"
)

// IntToHex는 int64를 바이트 배열로 변환합니다.
// proofofwork.go 파일에서 사용!
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}
