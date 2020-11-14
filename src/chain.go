package src

type Chain interface {
	NewBlock(proof string, previousHash string) *block
	LastBlock() *block
}

type chain struct {
	Blocks []*block
}

//	Create new empty chain with no blocks and transactions
func NewChain() *chain {
	return &chain{
		Blocks: make([]*block, 0),
	}
}

//	Add new block to the chain
func (c *chain) NewBlock(proof string, previousHash string) *block {
	return nil
}

//	return last block of the chain
func (c *chain) LastBlock() *block {
	return c.Blocks[len(c.Blocks)-1]
}
