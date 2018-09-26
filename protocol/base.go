package protocol

import "math/rand"

// DataGeneration is an interface that generate data
type DataGeneration interface {
	JustifyData()
	GenRandomData()
	SampleData()
}

// StructTypes are currently supported commands by bitmarkd
type StructTypes int

// lists all currently bitmarkd supported types
const (
	AssetsGetType StructTypes = iota
	BitmarkProvenanceType
	BitmarkTransferType
	BitmarksCreateType
	BitmarksProofType
	BlockOwnerTransferType
	NodeInfoType
	TransactionStatusType
	RandomType
)

// New create struct by StructTypes
func New(structName StructTypes) DataGeneration {
	var rpc DataGeneration
	var name StructTypes

	// if given nothing, use random one
	if structName == RandomType {
		name = StructTypes(rand.Intn(int(RandomType)))
	} else {
		name = structName
	}

	switch name {
	case AssetsGetType:
		rpc = &AssetsGetRpc{}
	case BitmarkProvenanceType:
		rpc = &BitmarkProvenanceRpc{}
	case BitmarkTransferType:
		rpc = &BitmarkTransferRpc{}
	case BitmarksCreateType:
		rpc = &BitmarksCreateRpc{}
	case BitmarksProofType:
		rpc = &BitmarksProofRpc{}
	case BlockOwnerTransferType:
		rpc = &BlockOwnerTransferRpc{}
	case NodeInfoType:
		rpc = &NodeInfoRpc{}
	case TransactionStatusType:
		rpc = &TransactionStatusRpc{}
	default:
	}

	return rpc
}
