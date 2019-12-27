package handlers

import (
	"github.com/algorand/go-algorand/plugins/rl/api/spec/v1"
)

// InfoResponse contains information about the RandLabs Extension
//
// swagger:response InfoResponse
type InfoResponse struct {
	// in: body
	Body *v1.RLInfo
}

func (sr InfoResponse) getBody() interface{} {
	return sr.Body
}

// DisassembleTealResponse contains the DisassembledTeal program details
type DisassembleTealResponse struct {
	Body *v1.DisassembledTeal
}

func (r DisassembleTealResponse) getBody() interface{} {
	return r.Body
}
