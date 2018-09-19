package protocol

import "local/random"

type NodeInfoRpc struct {
	ID     string     `json:"id,omitempty"`
	Method string     `json:"method,omitempty"`
	Params []NodeInfo `json:"params,omitempty"`
}

type NodeInfo struct{}

func (rpc *NodeInfoRpc) JustifyData() {
	rpc.Method = "Node.Info"
	rpc.ID = "1"
}

// genRandomData generates random data fits specific interface
func (rpc *NodeInfoRpc) GenRandomData() {
	r := random.New()
	r.Fuzz(rpc)
}

// sampleData generates correct data
func (rpc *NodeInfoRpc) SampleData() {
	rpc.ID = "1"
	rpc.Method = "Node.Info"
	rpc.Params = []NodeInfo{}
}
