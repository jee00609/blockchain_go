package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// 블록의 구조체
type Block struct {
	Timestamp     int64  //현재 시간의 타임스탬프 (블록 생성 시간)
	Data          []byte //블록에 포함된 실제 가치를 지닌 정보
	PrevBlockHash []byte //이전 블록의 해시값
	Hash          []byte //블록의 해시값
	Nonce         int    //(추가) 증명 검증
}

// SetHash 메서드를 작성
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

// 새로운 블럭 생성 및 블럭 리턴
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// 새 블록을 추가하려면 기존 블록이 필요하다.
// 그러나 우리 블록체인에는 아무 블록도 없다!
// 따라서 모든 블록체인에는 적어도 하나의 블록이 있어야하며, 블록체인의 첫번째 블록을 제네시스 블록이라고 한다.
// 이 블록을 생성하는 메서드를 구현한다.
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
