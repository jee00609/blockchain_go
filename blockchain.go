package main

type Blockchain struct{
	blocks []*Block
}

// 블록 추가 기능
func (bc *Blockchain) AddBlock(data string){
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data,prevBlock.Hash)
	bc.blocks = append(bc.blocks,newBlock)
}

// 제네시스 블록을 가지고 블록체인을 생성하는 함수
func NewBlockchain() *Blockchain{
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}