package protocol

import fuzz "github.com/jamieabc/gofuzz"

const (
	bitmarksCreateRPCMethod string = "Bitmarks.Create"
)

type BitmarksCreateRpc struct {
	ID     string           `json:"id"`
	Method string           `json:"method"`
	Params []BitmarksCreate `json:"params"`
}

type BitmarksCreate struct {
	Assets []AssetData `json:"assets"`
	Issues []IssueData `json:"issues"`
}

// AssetData reference form bitmarkd/transactionrecord/transaction.go, AssetData
type AssetData struct {
	Name        string `json:"name"`        // utf-8
	Fingerprint string `json:"fingerprint"` // utf-8
	Metadata    string `json:"metadata"`    // utf-8
	Registrant  string `json:"registrant"`  // base58
	Signature   string `json:"signature"`   // hex
}

// BitmarkIssue the unpacked BitmarkIssue structure
// AssetData reference form bitmarkd/transactionrecord/transaction.go, AssetData
type IssueData struct {
	AssetID   string `json:"assetId"`   // link to asset record
	Owner     string `json:"owner"`     // base58: the "destination" owner
	Nonce     uint64 `json:"nonce"`     // to allow for multiple issues at the same time
	Signature string `json:"signature"` // hex: corresponds to owner in linked record
}

func (rpc *BitmarksCreateRpc) JustifyData() {
	rpc.Method = bitmarksCreateRPCMethod
	rpc.ID = "1"
}

// GenRandomData generates random data fits specific interface
func (rpc *BitmarksCreateRpc) GenRandomData() {
	f := fuzz.New()
	f.Fuzz(rpc)
}

// SampleData generates correct data
func (rpc *BitmarksCreateRpc) SampleData() {
	rpc.ID = "1"
	rpc.Method = bitmarksCreateRPCMethod
	rpc.Params = []BitmarksCreate{
		BitmarksCreate{
			Assets: []AssetData{
				AssetData{
					Name:        "name",
					Fingerprint: "01840006653e9ac9e95117a15c915caab81662918e925de9e004f774ff82d7079a40d4d27b1b372657c61d46d470304c88c788b3a4527ad074d1dccbee5dbaa99a",
					Metadata:    "k1\u0000v1\u0000k2\u0000v2",
					Registrant:  "e1pFRPqPhY2gpgJTpCiwXDnVeouY9EjHY6STtKwdN6Z4bp4sog",
					Signature:   "dc9ad2f4948d5f5defaf9043098cd2f3c245b092f0d0c2fc9744fab1835cfb1ad533ee0ff2a72d1cdd7a69f8ba6e95013fc517d5d4a16ca1b0036b1f3055270c",
				},
			},
			Issues: []IssueData{
				IssueData{
					AssetID:   "3c50d70e0fe78819e7755687003483523852ee6ecc59fe40a4e70e89496c4d45313c6d76141bc322ba56ad3f7cd9c906b951791208281ddba3ebb5e7ad83436c",
					Owner:     "Owner:e1pFRPqPhY2gpgJTpCiwXDnVeouY9EjHY6STtKwdN6Z4bp4sog",
					Nonce:     4,
					Signature: "6ecf1e6d965e4364321596b4675950554b3b8f1b40be3deb64306ddf72fef09f3c6bcebd6375925a51b984f56ec751a54c88f0dab56b3f69708a7b634c428a0a",
				},
			},
		},
	}
}
