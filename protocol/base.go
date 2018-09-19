package protocol

type DataGeneration interface {
	JustifyData()
	GenRandomData()
	SampleData()
}

type StructTypes int

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

func New(structName StructTypes) DataGeneration {
	var rpc DataGeneration

	switch structName {
	case AssetsGetType:
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
