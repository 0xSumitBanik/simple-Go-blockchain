package handlers

import (
	"fmt"
	"go-blockchain/core"
)

// Container will hold all dependencies for your application.
type Container struct {
	Block      core.Block
	Blockchain core.Blockchain
}

// NewContainer returns an empty or an initialized container for your handlers.
func NewContainer() (Container, error) {
	c := Container{}
	fmt.Printf("Type: %T\n", c.Blockchain)
	return c, nil
}
