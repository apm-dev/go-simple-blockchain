package bc

import (
	"errors"
	"github.com/apm-dev/go-simple-blockchain/src/crypto"
)

type Chain interface {
	addBlock(block) error
	lastBlock() *block
	getBlocks() []*block
}

type chain struct {
	Blocks []*block
}

//	create new chain with genesis block
func newChain(gb block) Chain {
	c := chain{}
	c.Blocks = append(c.Blocks, &gb)
	return &c
}

//	Add new block to the chain
func (c *chain) addBlock(b block) error {
	lbh, err := crypto.Hash(crypto.HashSHA256, c.lastBlock())
	if err != nil {
		return err
	}
	if lbh != b.PreviousHash {
		return errors.New("the block is invalid")
	}
	c.Blocks = append(c.Blocks, &b)
	return nil
}

//	return last block of the chain
func (c *chain) lastBlock() *block {
	return c.Blocks[len(c.Blocks)-1]
}

func (c *chain) getBlocks() []*block {
	return c.Blocks
}
