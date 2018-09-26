package protocol

import "github.com/jamieabc/gofuzz"

const (
	assetsGetRPCMethod string = "Assets.Get"
)

type AssetsGetRpc struct {
	ID     string      `json:"id"`
	Method string      `json:"method"`
	Params []AssetsGet `json:"params"`
}

type AssetsGet struct {
	FingerPrints []string `json:"fingerprints"`
}

func (rpc *AssetsGetRpc) JustifyData() {
	rpc.Method = assetsGetRPCMethod
	rpc.ID = "1"
}

// GenRandomData generates random data fits specific interface
func (rpc *AssetsGetRpc) GenRandomData() {
	f := fuzz.New()
	f.Fuzz(rpc)
}

// SampleData generates correct data
func (rpc *AssetsGetRpc) SampleData() {
	rpc.ID = "1"
	rpc.Method = assetsGetRPCMethod
	rpc.Params = []AssetsGet{
		AssetsGet{
			FingerPrints: []string{
				"015b9c9fb3e993bf64500977844104fc0b70ef5ca99141e9f56e2e837ce668fbe787643c34a0d51a32a82408eb36a6e93f7badbc5af50de29d9401b5affe564440",
				"e12571eff187f9a88e8a516c639fc51ef6ff9472fd39fd326cec33e266ae9ce74863f428f1e153f724b19b4b1d26df586f1ea3b794a5ca617b37129d315e3918"},
		},
	}
}
