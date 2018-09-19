package protocol

import "local/random"

type AssetsGetRpc struct {
	ID     string      `json:"id,omitempty"`
	Method string      `json:"method,omitempty"`
	Params []AssetsGet `json:"params,omitempty"`
}

type AssetsGet struct {
	Count uint   `json:"count,omitempty"`
	TxID  string `json:"txId,omitempty"`
}

func (rpc *AssetsGetRpc) JustifyData(methodName string) {
	rpc.Method = methodName
	rpc.ID = "1"
	for idx := range rpc.Params {
		rpc.Params[idx].Count = 20
	}
}

// genRandomData generates random data fits specific interface
func (rpc *AssetsGetRpc) GenRandomData() {
	r := random.New()
	r.Fuzz(rpc)
}

// sampleData generates correct data
func (rpc *AssetsGetRpc) SampleData() {
	rpc.ID = "1"
	rpc.Method = "Bitmark.Provenance"
	rpc.Params = []AssetsGet{
		AssetsGet{
			Count: 20,
			TxID:  "2dc8770718b01f0205ad991bfb4c052f02677cff60e65d596e890cb6ed82c861",
		},
	}
}
