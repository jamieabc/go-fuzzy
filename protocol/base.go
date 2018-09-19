package protocol

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
)

// create struct by StructTypes
func New(structName StructTypes) DataGeneration {
	var rpc DataGeneration

	switch structName {
	case AssetsGetType:
		rpc = &AssetsGetRpc{}
	case BitmarkProvenanceType:
		rpc = &BitmarkProvenanceRpc{}
	case BitmarkTransferType:
	case BitmarksCreateType:
	case BitmarksProofType:
		rpc = &BitmarksProofRpc{}
	case BlockOwnerTransferType:
	case NodeInfoType:
		rpc = &NodeInfoRpc{}
	case TransactionStatusType:
		rpc = &TransactionStatusRpc{}
	default:
	}

	return rpc
}
