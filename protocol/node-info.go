package protocol

import fuzz "github.com/jamieabc/gofuzz"

const (
	nodeInfoRPCMethod string = "Node.Info"
)

type NodeInfoRpc struct {
	ID     string     `json:"id"`
	Method string     `json:"method"`
	Params []NodeInfo `json:"params"`
}

type NodeInfo struct{}

func (rpc *NodeInfoRpc) JustifyData() {
	rpc.Method = nodeInfoRPCMethod
	rpc.ID = "1"
}

// GenRandomData generates random data fits specific interface
func (rpc *NodeInfoRpc) GenRandomData() {
	f := fuzz.New()
	f.Fuzz(rpc)
}

// SampleData generates correct data
func (rpc *NodeInfoRpc) SampleData() {
	rpc.ID = "1"
	rpc.Method = nodeInfoRPCMethod
	rpc.Params = []NodeInfo{}
}
