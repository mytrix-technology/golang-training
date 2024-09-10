package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type BlockData struct {
	Location   string
	WaveHeight int
}

type Block struct {
	Data         BlockData
	Hash         string
	PreviousHash string
	Timestamp    int64
	Height       int
	Pow          int
}

func (b Block) calculateHash() string {
	data, _ := json.Marshal(b.Data)
	blockData := b.PreviousHash + string(data) + strconv.FormatInt(b.Timestamp, 10) + strconv.Itoa(b.Height) + strconv.Itoa(b.Pow)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

func (b *Block) mine(difficulty int) {
	for !strings.HasPrefix(b.Hash, strings.Repeat("0", difficulty)) {
		b.Pow++
		b.Hash = b.calculateHash()
	}
}
