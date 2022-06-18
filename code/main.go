package main

import (
	"fmt"
	"time"
	"crypto/sha256"
	"strconv"
	"bytes"
)


type Block struct{
	Timestamp int64
	Data []byte
	PrevBlockHash []byte
	Hash []byte
}

type Blockchain struct {
	blocks []*Block
}


func main(){
	// blockchainの先頭block生成
	bc := NewBlockChain()
	// ポインタなので中身見るならこんな書き方
    // fmt.Printf("%+v\n", *bc.blocks[0])
	// {Timestamp:1655567100 Data:[71 101 110 101 115 105 115 32 66 108 111 99 107] PrevBlockHash:[] Hash:[105 120 133 246 222 58 61 199 200 245 126 23 191 25 138 43 6 218 84 48 150 2 124 176 144 120 218 89 172 231 171 219]}

	// 後続のblock生成
	bc.AddBlock("Send 1 BTC To Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")
	bc.AddBlock("Send 3 more BTC to Ivan")// tekitouni tuika
	bc.AddBlock("Send 4 more BTC to Ivan")// tekitouni tuika
	bc.AddBlock("Send 5 more BTC to Ivan")// tekitouni tuika
	bc.AddBlock("Send 6 more BTC to Ivan")// tekitouni tuika

	for _, block := range bc.blocks{
		fmt.Printf("Prev. Hash = %x\n", block.PrevBlockHash)
		fmt.Printf("Data = %s\n", block.Data)
		fmt.Printf("Hash = %x\n", block.Hash)
		fmt.Println()
	}
}

func (b *Block) SetHash(){
	// 実行時刻のUnix timeを10進数の文字列へ
	timestamp :=[]byte(strconv.FormatInt(b.Timestamp, 10))
	// もってるデータをbyte列で連結
	// ex) [[144 233 241] [83 101 110] [49 54 53]] => [144 233 241 83 101 110 49 54 53]
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	// 連結したデータからハッシュ値計算してセット
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func NewBlock(data string, PrevBlockHash []byte) *Block{
	block := &Block{
		time.Now().Unix(),
		[]byte(data),
		PrevBlockHash,
		[]byte{}, // Hash部分を空で用意, かな. この書き方だとアロケートしないんだっけ.
	}
	// hashは別で計算
	block.SetHash()

	return block
}


func (bc *Blockchain) AddBlock(data string){
	// ひとつ前のblock情報取ってきて次のblockを生成
	prevBlock :=bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}


func NewGenesisBlock() *Block{
	// 先頭block生成時は PrevBlockHash はないので空のbyteを渡す
	return NewBlock("Genesis Block", []byte{})
}

func NewBlockChain() *Blockchain{
	// Blockchain型の配列に入れて返す
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}