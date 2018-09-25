package protocol

import "local/random"

const (
	bitmarksProofRPCMethod string = "Bitmarks.Proof"
)

type BitmarksProofRpc struct {
	ID     string          `json:"id,omitempty"`
	Method string          `json:"method,omitempty"`
	Params []BitmarksProof `json:"params,omitempty"`
}

type BitmarksProof struct {
	PayId string `json:"payId"`
	Nonce string `json:"nonce"`
}

func (rpc *BitmarksProofRpc) JustifyData() {
	rpc.Method = bitmarksProofRPCMethod
	rpc.ID = "1"
}

// genRandomData generates random data fits specific interface
func (rpc *BitmarksProofRpc) GenRandomData() {
	r := random.New()
	r.Fuzz(rpc)
}

// sampleData generates correct data
func (rpc *BitmarksProofRpc) SampleData() {
	rpc.ID = "1"
	rpc.Method = bitmarksProofRPCMethod
	rpc.Params = []BitmarksProof{
		BitmarksProof{
			PayId: "e219ffe8021190ea472baa147b05b2fbfb79818ab6eb267a037b7744f3b89a5966723180e1367ce7e172369d7432a658",
			Nonce: "c114fa516a98c3de",
		},
	}
}
