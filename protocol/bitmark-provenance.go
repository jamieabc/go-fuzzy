package protocol

import "local/random"

const (
	bitmarkProvenanceRPCMethod string = "Bitmark.Provenance"
)

type BitmarkProvenanceRpc struct {
	ID     string       `json:"id,omitempty"`
	Method string       `json:"method,omitempty"`
	Params []Provenance `json:"params,omitempty"`
}

type Provenance struct {
	Count uint   `json:"count,omitempty"`
	TxID  string `json:"txId,omitempty"`
}

func (rpc *BitmarkProvenanceRpc) JustifyData() {
	rpc.Method = bitmarkProvenanceRPCMethod
	rpc.ID = "1"
	for idx := range rpc.Params {
		rpc.Params[idx].Count = 20
	}
}

// genRandomData generates random data fits specific interface
func (rpc *BitmarkProvenanceRpc) GenRandomData() {
	r := random.New()
	r.Fuzz(rpc)
}

// sampleData generates correct data
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
