package protocol

import fuzz "github.com/jamieabc/gofuzz"

const (
	bitmarkProvenanceRPCMethod string = "Bitmark.Provenance"
)

type BitmarkProvenanceRpc struct {
	ID     string       `json:"id"`
	Method string       `json:"method"`
	Params []Provenance `json:"params"`
}

type Provenance struct {
	Count uint   `json:"count"`
	TxID  string `json:"txId"`
}

func (rpc *BitmarkProvenanceRpc) JustifyData() {
	rpc.Method = bitmarkProvenanceRPCMethod
	rpc.ID = "1"
	for idx := range rpc.Params {
		rpc.Params[idx].Count = 20
	}
}

// GenRandomData generates random data fits specific interface
func (rpc *BitmarkProvenanceRpc) GenRandomData() {
	f := fuzz.New()
	f.Fuzz(rpc)
}

// SampleData generates correct data
func (rpc *BitmarkProvenanceRpc) SampleData() {
	rpc.ID = "1"
	rpc.Method = bitmarkProvenanceRPCMethod
	rpc.Params = []Provenance{
		Provenance{
			Count: 20,
			TxID:  "2dc8770718b01f0205ad991bfb4c052f02677cff60e65d596e890cb6ed82c861",
		},
	}
}
