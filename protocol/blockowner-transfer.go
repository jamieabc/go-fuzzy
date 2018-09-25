package protocol

import (
	"local/random"
)

const (
	blockownerTransferRPCMethod string = "BlockOwner.Transfer"
)

type BlockOwnerTransferRpc struct {
	ID     string               `json:"id,omitempty"`
	Method string               `json:"method,omitempty"`
	Params []BlockOwnerTransfer `json:"params,omitempty"`
}

type Currency uint64

type CurrencyMapping map[Currency]string

type BlockOwnerTransfer struct {
	Link             string           `json:"link"`             // previous record
	Escrow           *TransferPayment `json:"escrow"`           // optional escrow TransferPayment address, looks like a third party voucher to ensure payment will be paid
	Version          uint64           `json:"version"`          // reflects combination of supported currencies
	Payments         *CurrencyMapping `json:"payments"`         // require length and contents depend on version
	Owner            string           `json:"owner"`            // base58: the "destination" owner
	Signature        string           `json:"signature"`        // hex: corresponds to owner in linked record
	Countersignature string           `json:"countersignature"` // hex: corresponds to owner in this record, used in two signature
}

func (rpc *BlockOwnerTransferRpc) JustifyData() {
	rpc.Method = blockownerTransferRPCMethod
	rpc.ID = "1"
}

// genRandomData generates random data fits specific interface
func (rpc *BlockOwnerTransferRpc) GenRandomData() {
	r := random.New()
	r.Fuzz(rpc)
}

// sampleData generates correct data
func (rpc *BlockOwnerTransferRpc) SampleData() {
	rpc.ID = "1"
	rpc.Method = blockownerTransferRPCMethod
	rpc.Params = []BlockOwnerTransfer{
		BlockOwnerTransfer{
			Link:    "1bebd06c8ecb8b11ea93e93c9d38b7f6d7dfdf015530819015172cf51c7f33f7",
			Version: 5,
			Payments: &CurrencyMapping{
				1: "BTC",
			},
			Owner:     "eZpG6Wi9SQvpDatEP7QGrx6nvzwd6s6R8DgMKgDbDY1R5bjzb9",
			Signature: "a3e456a31a4a64962a32bcbf6549d14134deeb5d87285a04c648355eb9e59d938f8ab440d2b50c781baf2c1a5a2112c2167301bb128c8f850a9d54f3b27c5a08",
		},
	}
}
