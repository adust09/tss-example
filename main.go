package main

import (
	"math/big"
	"time"

	"github.com/bnb-chain/tss-lib/common"
	"github.com/bnb-chain/tss-lib/ecdsa/keygen"
	"github.com/bnb-chain/tss-lib/ecdsa/signing"
	"github.com/bnb-chain/tss-lib/tss"
)

func Main() {
	ids := []tss.UnSortedPartyIDs{}
	parties := tss.SortPartyIDs(ids[0], 0)
	partyIDMap := make(map[string]*tss.PartyID)
	for _, id := range parties {
		partyIDMap[id.Id] = id
	}

	curve := tss.S256()
	ctx := tss.NewPeerContext(parties)
	thisParty := tss.NewPartyID("Bitcoin", "BTC", big.NewInt(1))
	params := tss.NewParameters(curve, ctx, thisParty, len(parties), 3)

	KeyGen(params)
	Sign(params)
}

func KeyGen(params *tss.Parameters) {
	outCh := make(chan tss.Message, 100)
	endCh := make(chan keygen.LocalPartySaveData, 1)
	preParams, _ := keygen.GeneratePreParams(1 * time.Minute)

	party := keygen.NewLocalParty(params, outCh, endCh, *preParams)

	party.Start()

}

func Sign(params *tss.Parameters) {
	keys, _, _ := keygen.LoadKeygenTestFixturesRandomSet(3, 10)

	outCh := make(chan tss.Message, 100)
	endCh := make(chan common.SignatureData, 1)

	party := signing.NewLocalParty(big.NewInt(42), params, keys[0], outCh, endCh)
	party.Start()
}
