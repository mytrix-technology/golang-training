package main

//type Blockchain struct {
//	GenesisBlock Block
//	Chain        []Block
//	Difficulty   int
//}
//
//func NewBlockchain(difficulty int) Blockchain {
//	genesisBlock := Block{
//		Hash:      "0",
//		Height:    0,
//		Timestamp: time.Now().Unix(),
//	}
//	return Blockchain{
//		genesisBlock,
//		[]Block{genesisBlock},
//		difficulty,
//	}
//}
//
//func (b *Blockchain) appendBlock(location string, waveHeight int) {
//	blockData := BlockData{
//		Location:   location,
//		WaveHeight: waveHeight,
//	}
//	lastBlock := b.Chain[len(b.Chain)-1]
//	newBlock := Block{
//		Data:         blockData,
//		PreviousHash: lastBlock.Hash,
//		Timestamp:    time.Now().Unix(),
//		Height:       lastBlock.Height + 1,
//	}
//	newBlock.mine(b.Difficulty)
//	b.Chain = append(b.Chain, newBlock)
//}
//
//func (b Blockchain) isValid() bool {
//	for i := range b.Chain[1:] {
//		previousBlock := b.Chain[i]
//		currentBlock := b.Chain[i+1]
//		if currentBlock.Height != previousBlock.Height+1 {
//			log.Println("Bad Height")
//			return false
//		}
//		if currentBlock.Hash != currentBlock.calculateHash() {
//			log.Println("Bad Hash")
//			return false
//		}
//		if currentBlock.PreviousHash != previousBlock.Hash {
//			log.Println("Bad Prev Hash")
//			return false
//		}
//	}
//	return true
//}
