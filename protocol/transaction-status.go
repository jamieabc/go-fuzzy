package protocol

import fuzz "github.com/jamieabc/gofuzz"

const (
	transactionStatusRPCMethod string = "Transaction.Status"
)

type TransactionStatusRpc struct {
	ID     string              `json:"id"`
	Method string              `json:"method"`
	Params []TransactionStatus `json:"params"`
}

type TransactionStatus struct {
	TxID string `json:"TxId"`
}

func (rpc *TransactionStatusRpc) JustifyData() {
	rpc.Method = transactionStatusRPCMethod
	rpc.ID = "1"
}

// GenRandomData generates random data fits specific interface
func (rpc *TransactionStatusRpc) GenRandomData() {
	f := fuzz.New()
	f.Fuzz(rpc)
}

// SampleData generates correct data
func (rpc *TransactionStatusRpc) SampleData() {
	rpc.ID = "1"
	rpc.Method = transactionStatusRPCMethod
	rpc.Params = []TransactionStatus{
		TransactionStatus{
			TxID: "2dc8770718b01f0205ad991bfb4c052f02677cff60e65d596e890cb6ed82c861",
		},
	}
}
