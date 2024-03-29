package blockchain

type Blockchain struct {
	Blocks []*Block
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func (bc Blockchain) Validate() bool {
	for _, block := range bc.Blocks {
		pow := NewProofOfWork(block)
		if !pow.Validate() {
			return false
		}
	}
	return true
}
