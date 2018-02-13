package btcinsightapiclient

import (
	"strconv"

	"github.com/tidwall/gjson"
)

// LatestHeight gets latest height
func LatestHeight(apiURL string) (int64, error) {
	bs, err := invokeHTTP(apiURL + "/blocks?limit=1")
	if err != nil {
		return 0, err
	}
	return gjson.GetBytes(bs, "blocks.0.height").Int(), nil
}

// TxOut type
type TxOut struct {
	TxID        string
	BlockHeight int64
	N           int
	Value       string
	Address     string
}

// SyncHeight syncs block at given height
func SyncHeight(apiURL string, height int64) ([]*TxOut, error) {
	hash, err := fetchBlockHash(apiURL, height)
	if err != nil {
		return nil, err
	}

	txs, err := fetchTxList(apiURL, hash)
	if err != nil {
		return nil, err
	}

	xs := make([]*TxOut, 0)

	for _, tx := range txs {
		t, err := fetchTx(apiURL, tx)
		if err != nil {
			return nil, err
		}
		xs = append(xs, t...)
	}

	return xs, nil
}

func fetchBlockHash(apiURL string, height int64) (string, error) {
	bs, err := invokeHTTP(apiURL + "/block-index/" + strconv.FormatInt(height, 10))
	if err != nil {
		return "", err
	}
	return gjson.GetBytes(bs, "blockHash").String(), nil
}

func fetchTxList(apiURL string, blockHash string) ([]string, error) {
	bs, err := invokeHTTP(apiURL + "/block/" + blockHash)
	if err != nil {
		return nil, err
	}
	r := gjson.GetBytes(bs, "tx").Array()
	xs := make([]string, len(r))
	for i := range r {
		xs[i] = r[i].String()
	}
	return xs, nil
}

func fetchTx(apiURL string, txID string) ([]*TxOut, error) {
	bs, err := invokeHTTP(apiURL + "/tx/" + txID)
	if err != nil {
		return nil, err
	}

	blockHeight := gjson.GetBytes(bs, "blockheight").Int()
	txID = gjson.GetBytes(bs, "txid").String()

	vouts := gjson.GetBytes(bs, "vout").Array()
	xs := make([]*TxOut, 0, len(vouts))
	for _, vout := range vouts {
		x := TxOut{
			BlockHeight: blockHeight,
			TxID:        txID,
			Address:     vout.Get("scriptPubKey.addresses.0").String(),
			N:           int(vout.Get("n").Int()),
			Value:       vout.Get("value").String(),
		}
		if x.Address == "" {
			continue
		}
		xs = append(xs, &x)
	}
	return xs, nil
}
