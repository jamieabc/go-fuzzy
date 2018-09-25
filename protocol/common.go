package protocol

type TransferPayment struct {
	Currency uint64 `json:"currency"`      // utf-8 → Enum
	Address  string `json:"address"`       // utf-8
	Amount   uint64 `json:"amount,string"` // number as string, in terms of smallest currency unit
}
