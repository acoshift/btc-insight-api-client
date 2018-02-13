package btcinsightapiclient_test

import (
	"fmt"
	"testing"

	btcinsightapiclient "github.com/acoshift/btc-insight-api-client"
)

const apiURL = "https://cashexplorer.bitcoin.com/api"

func TestSyncHeight(t *testing.T) {
	txs, err := btcinsightapiclient.SyncHeight(apiURL, 517200)
	if err != nil {
		t.Fatal(err)
	}
	for _, tx := range txs {
		fmt.Println(tx.BlockHeight, tx.TxID, tx.N, tx.Address, tx.Value)
	}
}

func TestLatestHeight(t *testing.T) {
	height, err := btcinsightapiclient.LatestHeight(apiURL)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(height)
}
