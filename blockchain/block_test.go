package blockchain

import (
	"strconv"
	"testing"

	"go-demo/blockchain/core"
)

func TestBlock(t *testing.T) {
	t.Log(core.GenerateGenesisBlock())
}

func TestBlockChain(t *testing.T) {
	chain := core.NewBlockChain()
	for i := 0; i < 10; i++ {
		chain.SendData("block:" + strconv.Itoa(i))
	}

	for _, value := range chain.Blocks {
		t.Logf("%+v", value)
	}
}
