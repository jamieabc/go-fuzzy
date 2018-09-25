package protocol

import "local/random"

const (
	transactionStatusRPCMethod string = "Transaction.Status"
)

type TransactionStatusRpc struct {
	ID     string              `json:"id,omitempty"`
	Method string              `json:"method,omitempty"`
	Params []TransactionStatus `json:"params,omitempty"`
}

type TransactionStatus struct {
	TxID string `json:"TxId,omitempty"`
}

func (rpc *TransactionStatusRpc) JustifyData() {
	rpc.Method = transactionStatusRPCMethod
	rpc.ID = "1"
}

// genRandomData generates random data fits specific interface
func (rpc *TransactionStatusRpc) GenRandomData() {
	r := random.New()
	r.Fuzz(rpc)
}

// sampleData generates correct data
func (rpc *TransactionStatusRpc) SampleData() {
	rpc.ID = "1"
	rpc.Method = transactionStatusRPCMethod
	rpc.Params = []TransactionStatus{
		TransactionStatus{
			TxID: "2dc8770718b01f0205ad991bfb4c052f02677cff60e65d596e890cb6ed82c861",
		},
	}
}
