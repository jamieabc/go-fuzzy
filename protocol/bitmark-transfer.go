package protocol

import fuzz "github.com/jamieabc/gofuzz"

const (
	bitmarkTransferRPCMethod string = "Bitmark.Transfer"
)

type BitmarkTransferRpc struct {
	ID     string            `json:"id"`
	Method string            `json:"method"`
	Params []BitmarkTransfer `json:"params"`
}

type BitmarkTransfer struct {
	Link             string           `json:"link"`             // previous record
	Escrow           *TransferPayment `json:"escrow"`           // optional escrow TransferPayment address, looks like a third party voucher to ensure payment will be paid
	Owner            string           `json:"owner"`            // base58: the "destination" owner
	Signature        string           `json:"signature"`        // hex: corresponds to owner in linked record
	Countersignature string           `json:"countersignature"` // hex: corresponds to owner in this record, used in two signature
}

func (rpc *BitmarkTransferRpc) JustifyData() {
	rpc.Method = bitmarkTransferRPCMethod
	rpc.ID = "1"
}

// GenRandomData generates random data fits specific interface
func (rpc *BitmarkTransferRpc) GenRandomData() {
	f := fuzz.New()
	f.Fuzz(rpc)
}

// SampleData generates correct data
func (rpc *BitmarkTransferRpc) SampleData() {
	rpc.ID = "1"
	rpc.Method = bitmarkTransferRPCMethod
	rpc.Params = []BitmarkTransfer{
		BitmarkTransfer{
			Link:      "1bebd06c8ecb8b11ea93e93c9d38b7f6d7dfdf015530819015172cf51c7f33f7",
			Owner:     "eZpG6Wi9SQvpDatEP7QGrx6nvzwd6s6R8DgMKgDbDY1R5bjzb9",
			Signature: "a3e456a31a4a64962a32bcbf6549d14134deeb5d87285a04c648355eb9e59d938f8ab440d2b50c781baf2c1a5a2112c2167301bb128c8f850a9d54f3b27c5a08",
		},
	}
}
