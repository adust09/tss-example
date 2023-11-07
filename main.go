package main

import (
	"fmt"
	"time"

	"github.com/bnb-chain/tss-lib/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/tss"
)

func main() {
	start := time.Now()

	pIDs := tss.GenerateTestPartyIDs(5)
	p2pCtx := tss.NewPeerContext(pIDs)
	threshold := 2
	params := tss.NewParameters(tss.S256(), p2pCtx, pIDs[0], len(pIDs), threshold)

	KeyGen(params)

	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("実行時間: %s\n", elapsed)

	// Sign(params)
}

func KeyGen(params *tss.Parameters) {
	outCh := make(chan tss.Message, 100)
	endCh := make(chan keygen.LocalPartySaveData, 1)
	preParams, _ := keygen.GeneratePreParams(1 * time.Minute)
	party := keygen.NewLocalParty(params, outCh, endCh, *preParams)

	party.Start()

}

// func Sign(params *tss.Parameters) {
// 	keys, signPIDs, _ := keygen.LoadKeygenTestFixturesRandomSet(testThreshold+1, testParticipants)

// 	outCh := make(chan tss.Message, 100)
// 	endCh := make(chan common.SignatureData, 1)

// 	party := signing.NewLocalParty(big.NewInt(42), params, keys[0], outCh, endCh)
// 	party.Start()
// }
