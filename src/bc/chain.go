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
func newChain() Chain {
	return &chain{}
}

//	Add new block to the chain
func (c *chain) addBlock(b block) error {
	//	Validate new block previous hash with last block hash
	//	if it was not genesis block
	if c.lastBlock() != nil {
		lbh, err := crypto.Hash(crypto.HashSHA256, c.lastBlock())
		if err != nil {
			return err
		}
		if lbh != b.PreviousHash {
			return errors.New("the block is invalid")
		}
	}
	c.Blocks = append(c.Blocks, &b)
	return nil
}

//	return last block of the chain
func (c *chain) lastBlock() *block {
	if len(c.Blocks) == 0 {
		return nil
	}
	return c.Blocks[len(c.Blocks)-1]
}

func (c *chain) getBlocks() []*block {
	return c.Blocks
}
