package v1

import "github.com/algorand/go-algorand/data/transactions/logic"

// RLInfo contains the information about a node status
// swagger:model RLInfo
type RLInfo struct {
	// Details provides some details about the RandLabs API Extension
	// Required: true
	Details string `json:"details"`
}

// DisassembledTeal represents a disassembled TEAL program
type DisassembledTeal struct {
	Code    string          `json:"code"`
	Symbols logic.DebugData `json:"symbols"`
}
