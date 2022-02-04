package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Hash         []byte
	Data         []byte
	PreviousHash []byte
}

type Blockchain struct {
	blocks []*Block
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PreviousHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, previousHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), previousHash}
	block.DeriveHash()
	return block
}

func (chain *Blockchain) AddBlock(data string) {
	previousBlock := chain.blocks[len(chain.blocks)-1]
	newBlock := CreateBlock(data, previousBlock.Hash)
	chain.blocks = append(chain.blocks, newBlock)

}

func GenesisBlock() *Block {
	return CreateBlock("Genesis", nil)
}



func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock()}}
}

func main() {
	chain := InitBlockchain()
	chain.AddBlock("This is a first block")
	chain.AddBlock("This is a second block")
	chain.AddBlock("This is a third block")


	for _ , block := range chain.blocks {
		fmt.Printf("Prev Hash: %#x\n", block.PreviousHash)
		fmt.Printf("Hash: %#x\n", block.Hash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Println()
	}
}