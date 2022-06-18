package main

import ("fmt")


type Block struct{
	Timestamp int64
	Data []byte
	PrevBlockHazsh []byte
	Hash []byte
}

func (b *Block) SetHash(){
	fmt.Println(strconv.FormatInt(b.Timestamp, 10))

	timestamp :=[]byte(strconv.FormatInt(b.Timestamp, 10))

	headers := bytes.Join([][]byte{b.PreBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

// issue と紐づけできてなかった

func main(){
	fmt.Println("test")
}