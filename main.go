package main

import (
	// "bytes"
	// "crypto/sha256"

	"fmt"
	"go-blockchain/handlers"
)

// func (c *Container) DeriveHash() {
// 	info := bytes.Join([][]byte{c.Block.Data, c.Block.PreviousHash}, []byte{})
// 	hash := sha256.Sum256(info)
// 	c.Block.Hash = hash[:]
// }

// func zyx(data string, previousHash []byte)

// func (c *Container) CreateBlock(data string, previousHash []byte) *gobc.Block {
// 	block := &c.Block{[]byte{}, []byte(data), previousHash}
// 	block.DeriveHash()
// 	return block
// }

// func (c *Container) AddBlock(data string) {
// 	previousBlock := c.Blockchain.blocks[len(chain.blocks)-1]
// 	newBlock := CreateBlock(data, previousBlock.Hash)
// 	chain.blocks = append(chain.blocks, newBlock)

// }

// func GenesisBlock() *Block {
// 	return CreateBlock("Genesis", nil)
// }

// func InitBlockchain() *gobc.Blockchain {
// 	return &Blockchain{[]*Block{GenesisBlock()}}
// }

func main() {
	c, _ := handlers.NewContainer()
	fmt.Printf("%v",c.Block.Data)
	// chain := InitBlockchain()
	// chain.AddBlock("This is a first block")
	// chain.AddBlock("This is a second block")
	// chain.AddBlock("This is a third block")

	// for _, block := range chain.blocks {
	// 	fmt.Printf("Prev Hash: %#x\n", block.PreviousHash)
	// 	fmt.Printf("Hash: %#x\n", block.Hash)
	// 	fmt.Printf("Data: %s\n", block.Data)
	// 	fmt.Println()
	// }
}
