package internal

type Blockchain struct {
	Blocks []*Block
}

func CreateGenesisBlock() *Block {
	return CreateBlock([]byte("創世區塊"), []byte{})
}

func CreateBlockchain() *Blockchain {
	return &Blockchain{[]*Block{CreateGenesisBlock()}}
}

func (blockchain *Blockchain) AddBlock(Data []byte) {
	prevBlock := blockchain.Blocks[len(blockchain.Blocks)-1]
	newBlock := CreateBlock(Data, prevBlock.Hash)
	blockchain.Blocks = append(blockchain.Blocks, newBlock)
}
